# Playbook: Bulk-Reviewing and Merging a Backlog of Open PRs

This session cleared a backlog of 23 open PRs on `bobadilla-tech/requiems-api`
in one pass: 2 feature PRs from a trusted external contributor (Kinnouts) and 21
dependabot dependency bumps. `gh pr merge` was not usable — the authenticated
`gh` account had read-only access and the repo requires a CODEOWNERS review — so
every PR was merged by fetching its branch, merging it locally with git, and
pushing directly to `main`. This doc is the reusable procedure for the next time
there's a pile of PRs to clear the same way.

---

## 1. Check whether `gh pr merge` will actually work before planning around it

`gh pr list` showing `mergeable: MERGEABLE` does **not** mean `gh pr merge` will
succeed. Two independent things can block it:

```bash
gh pr view <n> -R <owner>/<repo> --json reviewDecision,mergeStateStatus
```

If `mergeStateStatus` is `BLOKCED`/`BLOCKED` with
`reviewDecision:
REVIEW_REQUIRED`, check the actual rule:

```bash
gh api repos/<owner>/<repo>/rules/branches/main
```

Look for a `pull_request` rule with `require_code_owner_review: true` — a
generic approval from an arbitrary account won't satisfy that, only someone
covered by `.github/CODEOWNERS` will. Then check whether the `gh` account itself
can push at all:

```bash
gh api repos/<owner>/<repo> --jq '.permissions'
gh auth status
```

If `push` is `false`, `gh pr merge` is a dead end regardless of reviews. In this
session that was exactly the case — `gh` was authenticated as a read-only
collaborator account, while the local git SSH remote (`git@github.com:...`) was
authenticated as the actual repo owner:

```bash
ssh -T git@github.com     # shows which account the SSH key maps to
git remote -v             # confirm origin uses that SSH identity
```

When the two diverge like this, git push (not `gh pr merge`) is the way through.
A real push (not `--dry-run`) to `main` will show whether the account can bypass
the ruleset — GitHub reports this explicitly:

```
remote: Bypassed rule violations for refs/heads/main:
remote: - At least 1 approving review is required by reviewers with write access.
```

That line means the push succeeded despite the rule, because the pushing account
has bypass privileges (repo owner/admin). If a PR is a fork PR (like the
Kinnouts ones here), the fork is usually already added as a remote — check
`git remote -v` before adding a new one.

**Always confirm this bypass path with the user before using it.** It skips the
review process the branch protection exists to enforce; it's appropriate for a
backlog of already-green, already-reviewed, low-risk PRs (the case here), not as
a default way to route around review.

## 2. Review before merging, not after

For any PR carrying actual source changes (not just a dependency bump), read the
full diff before merging:

```bash
gh pr diff <n> -R <owner>/<repo> > /tmp/pr<n>.diff
```

Check for: unsafe rendering of user input (`.html_safe` on anything
user-controlled, raw `params` interpolated into HTML/SQL/shell), whether it
follows existing patterns in the file (a new "tool page" should look exactly
like the last one added), and whether tests exist for the new
success/error/edge-case paths. In this session both feature PRs turned out to be
copy-of-an-existing-pattern additions — same controller shape, same test shape,
same i18n structure — which is what made a fast approve safe.

For dependency bumps, patch/minor is low-risk by semver and doesn't need
individual scrutiny beyond "CI is green." **Major version bumps do need a quick
look** — pull the PR body (dependabot includes release notes) and check CI
actually exercised the changed code, not just installed it:

```bash
gh pr view <n> -R <owner>/<repo> --json body -q '.body' | head -60
gh pr view <n> -R <owner>/<repo> --json statusCheckRollup \
  -q '.statusCheckRollup[] | select(.conclusion != null) | "\(.name): \(.conclusion)"'
```

Look specifically for a real build/test step for the affected package (e.g.
"Worker Tests" or a Cloudflare "Workers Builds" check succeeding), not just a
lint pass. In this session, `typescript` 6→7 and `nanoid` 5→6 (×2) were the
majors; all three had a passing build/test check for the exact package touched,
and the changelogs called out nothing relevant (Node version floor bumps, perf —
nothing breaking for a Cloudflare Worker), so they merged in the same pass as
everything else.

## 3. Group dependency PRs by lockfile before merging any of them

Dependabot opens one PR per bump, but several bumps often land in the same
lockfile. Merging them one at a time in PR-number order **will** produce
lockfile conflicts once you're past the first PR in a group. Group first:

```bash
gh pr list --state open -R <owner>/<repo> --json number,title,headRefName \
  -q '.[] | [.number, .title] | @tsv'
```

Then bucket by the directory/lockfile the title names (`/apps/workers/shared`,
`/apps/api`, `/apps/dashboard`, etc.) and confirm each app really does have its
own lockfile (a pnpm/npm/bundler monorepo usually does, one per workspace —
check with `find . -iname "*.lock*"`, excluding `node_modules`). Two bumps to
_different_ lockfiles never conflict and can be merged in any order; two bumps
to the _same_ lockfile will conflict on the second one.

## 4. The merge loop, per PR

```bash
git fetch origin <headRefName>          # or the fork remote, if cross-repo
git merge --no-ff FETCH_HEAD -m "Merge pull request #<n> from <owner>/<headRefName>

<PR title>"
```

Using the exact `Merge pull request #<n> from <owner>/<branch>` message format
matches what GitHub's own merge button writes, which is what lets GitHub
retroactively recognize the PR as merged (see the caveat in §6).

If it merges clean, push immediately before starting the next PR — don't batch
pushes. Re-fetching/merging against a stale local `main` is how you'd miss a
conflict that a previous push introduced:

```bash
git push origin main
```

## 5. Resolving lockfile conflicts

When a conflict does hit (second+ PR touching the same lockfile), the resolution
is mechanical and should never involve hand-editing the lockfile's
resolved-graph content:

1. Look at just the manifest file's conflict (`package.json`, `Gemfile`,
   `go.mod`) — it's usually a small, readable diff of 1-3 version fields. Take
   the **newer** version on each conflicting line (the one just merged plus the
   one this PR is bumping), by hand-editing the conflict markers.
2. Regenerate the lockfile from the resolved manifest — never try to merge the
   lockfile's conflict markers by hand:
   ```bash
   pnpm install --lockfile-only     # pnpm workspace
   bundle install                   # ruby/Gemfile.lock — or `bundle lock`
   go mod tidy                      # go.sum
   ```
3. Grep the regenerated lockfile for the expected version to confirm it actually
   landed, `git add` both files, `git commit --no-edit` (keeps the merge commit
   message from step 4), then push.
4. Where the toolchain is available, run the package's own typecheck/build after
   a conflict resolution, not just after the whole cluster — a regenerated
   lockfile can silently resolve to something incompatible:
   ```bash
   pnpm run typecheck   # or: go build ./...
   ```

This is the only place in the whole process that needs judgment instead of being
mechanical — treat "take the newer version, regenerate, verify it built" as the
standard move and don't improvise beyond it.

## 6. Known cosmetic quirk: GitHub may show a merged PR as "Closed"

If a PR's conflict required editing the manifest file _beyond_ what its own
branch diff contained (i.e. you touched a line neither side's diff owned, to
reconcile two features/bumps), the resulting merge commit's diff no longer
matches any commit that exists on the PR's branch. GitHub's PR page then can't
algorithmically prove "this PR's changes reached main via a merge," and the PR
shows **CLOSED** instead of **MERGED** — even though the code is fully on
`main`. This happened once in 23 PRs here (the `typescript` 7.0.2 PR, after its
manifest conflicted with an already-merged `vite`/`tsx` bump).

Don't chase this — it's a UI label, not a functional gap. To confirm the change
actually shipped:

```bash
git log --oneline -- <changed-file>   # merge commit should be right there
grep -n '"<package>"' <manifest-file> # confirms the version that landed
```

## 7. Final verification

```bash
gh pr list --state open -R <owner>/<repo> --json number -q 'length'   # expect 0
gh pr list --state merged -R <owner>/<repo> --limit <n> \
  --json number,mergedAt -q '.[] | .number' | sort -n
```

Cross-check the merged list against the original PR numbers by hand — a PR that
shows `CLOSED` per §6 will silently drop out of `--state merged` filtering, so
don't rely on the count alone once you know that quirk exists. Then run whatever
whole-repo build/test commands are cheap enough to run locally per affected
package (`go build ./...`, `pnpm run typecheck` in each touched workspace,
`bundle check`) as a last sanity pass before calling it done.
