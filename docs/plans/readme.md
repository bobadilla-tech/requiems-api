# Plans

Record of AI-driven implementation plans, saved after execution as a changelog
of design decisions, why a change was made, not just what.

A plan starts as research: understand the current state related to the
problem/improvement/bug before deciding anything. Only after that does it become
an approach, what we're going to build/fix/test/benchmark, and how, described in
prose rather than code where possible (code only where the approach can't be
stated clearly without it). The write-up is finished with a notes section that's
explicitly meant to be filled in at the end of the work, not at planning time.

## Convention

- File name: `YYYY-MM-DD-short-topic.md`
- **Context** — the problem/improvement/bug, and the current state of the code
  relevant to it (what exists today, what's missing, prior attempts if any).
- **Approach** — what we're going to do and how, in prose over code. Cover
  build/fix/test/benchmark as applicable. Code snippets only when the approach
  genuinely can't be expressed without them.
- **Final notes** — filled in after the work lands, not while planning: what
  actually shipped, deviations from the approach, follow-ups. Every plan should
  have this section updated once the work is done.
- Save the plan after the work lands, not before — this is a record, not a
  backlog.
