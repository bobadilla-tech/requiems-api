#!/usr/bin/env bash
# tests/load/run.sh
#
# Convenience wrapper for the Requiem API k6 load-testing suite.
#
# Usage
# -----
#   ./tests/load/run.sh [scenario] [k6-options...]
#
# Scenarios
#   baseline          Baseline benchmark (1 VU, 2 min, all endpoints)
#   rate-limit        Rate limit enforcement validation
#   concurrent-users  Concurrent user simulation (ramp to 50 VUs)
#   all               Run all three scenarios in sequence (default)
#
# Examples
#   ./tests/load/run.sh
#   ./tests/load/run.sh baseline
#   ./tests/load/run.sh concurrent-users --out json=/tmp/results.json
#   BASE_URL=https://api.example.com ./tests/load/run.sh all
#   PEAK_VUS=100 ./tests/load/run.sh concurrent-users
#
# Requirements
#   k6 must be installed: https://k6.io/docs/get-started/installation/
#   The dev stack must be running: docker compose -f docker-compose.dev.yml up

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# ---------------------------------------------------------------------------
# Defaults
# ---------------------------------------------------------------------------

BASE_URL="${BASE_URL:-http://localhost:8080}"
SCENARIO="${1:-all}"
shift || true          # remaining args are forwarded to k6

# ---------------------------------------------------------------------------
# Validate k6 is installed
# ---------------------------------------------------------------------------

if ! command -v k6 &>/dev/null; then
  echo "ERROR: k6 is not installed."
  echo ""
  echo "Install instructions:"
  echo "  macOS:   brew install k6"
  echo "  Linux:   https://k6.io/docs/get-started/installation/"
  echo "  Docker:  docker run --rm -i grafana/k6 run -"
  exit 1
fi

echo "k6 $(k6 version)"
echo ""

# ---------------------------------------------------------------------------
# Runner
# ---------------------------------------------------------------------------

run_scenario() {
  local name="$1"
  local file="$SCRIPT_DIR/scenarios/${name}.ts"

  if [[ ! -f "$file" ]]; then
    echo "ERROR: Scenario file not found: $file"
    exit 1
  fi

  echo "=========================================="
  echo "  Running scenario: $name"
  echo "  Target:           $BASE_URL"
  echo "=========================================="
  echo ""

  BASE_URL="$BASE_URL" k6 run "$file" "$@"
  echo ""
}

# ---------------------------------------------------------------------------
# Dispatch
# ---------------------------------------------------------------------------

case "$SCENARIO" in
  baseline | rate-limit | concurrent-users)
    run_scenario "$SCENARIO" "$@"
    ;;
  all)
    run_scenario baseline "$@"
    run_scenario rate-limit "$@"
    run_scenario concurrent-users "$@"
    echo "All scenarios completed."
    ;;
  --help | -h | help)
    sed -n '/^# /p' "$0" | sed 's/^# //'
    ;;
  *)
    echo "ERROR: Unknown scenario '$SCENARIO'"
    echo "Valid scenarios: baseline, rate-limit, concurrent-users, all"
    exit 1
    ;;
esac
