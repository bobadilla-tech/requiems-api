#!/usr/bin/env bash

# Builds the body for the automated skill-regeneration PR: notes what
# triggered the run, pings the review team, and links the requiems.xyz
# doc page for each API that changed so reviewers can spot-check it.

set -euo pipefail

OUTPUT_FILE="${1:-pr-body.md}"

{
  echo "Skills regenerated automatically from the latest \`requiems-api\` documentation."
  echo ""
  echo "Triggered by: ${GITHUB_EVENT_NAME:-unknown}"
  echo ""
  echo "cc @bobadilla-tech/requiems-api — please review."

  apis="$(git status --porcelain -- skills/ | awk '{print $2}' | cut -d/ -f2 | sed -E 's/-(get|post|put|patch|delete)-.*$//' | sort -u)"
  if [ -n "$apis" ]; then
    echo ""
    echo "<details>"
    echo "<summary>Changed API doc pages to spot-check</summary>"
    echo ""
    echo "Verify each endpoint below is still current — links are built from the API's \`api_id\`, which doesn't always match the source YAML filename 1:1."
    echo ""
    echo "$apis" | while read -r api; do
      echo "- [$api](https://requiems.xyz/en/apis/$api)"
    done
    echo ""
    echo "</details>"
  fi
} > "$OUTPUT_FILE"
