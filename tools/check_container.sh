#!/usr/bin/env bash
# One bounded, loopback-only smoke for the production Dockerfile.
set -euo pipefail

image="bcm-local-smoke:latest"
if (( $# > 0 )); then image="$1"; fi
container=""
health_file="$(mktemp)"
cleanup() {
  rm -f "$health_file"
  if [[ -n "$container" ]]; then
    docker rm -f "$container" >/dev/null 2>&1 || true
  fi
}
trap cleanup EXIT

docker build --file Dockerfile --tag "$image" .
container="$(docker run -d --publish 127.0.0.1::8787 "$image")"
hostport="$(docker port "$container" 8787/tcp)"
port="$(printf '%s\n' "$hostport" | sed -n 's/^127\.0\.0\.1:\([0-9][0-9]*\)$/\1/p' | head -n 1)"
[[ "$port" =~ ^[0-9]+$ ]] || { echo 'No loopback port published' >&2; exit 1; }
base="http://127.0.0.1:$port"

ready=0
for ((i=0;i<30;i++)); do
  if curl --fail --silent --max-time 2 "$base/health" > "$health_file"; then
    ready=1
    break
  fi
  if [[ "$(docker inspect --format '{{.State.Running}}' "$container")" != "true" ]]; then
    echo 'BCM image exited before HTTP startup' >&2
    exit 1
  fi
  sleep 1
done
if (( !ready )); then
  echo 'BCM image never passed its synthetic-only health check' >&2
  exit 1
fi

HEALTH_FILE="$health_file" python3 - <<'PY'
import json, os
with open(os.environ["HEALTH_FILE"], encoding="utf-8") as f:
    data = json.load(f)
assert data == {"status": "ok", "data_scope": "synthetic_demo_only"}, data
PY

# Coordinates must not be resolved against the fictional DEMO provider.
response="$(curl --silent --show-error --max-time 5 --output /dev/null --write-out '%{http_code}' \
  -H 'Content-Type: application/json' \
  -d '{"project_id":"image-smoke","point":{"longitude":-105,"latitude":40},"applicability_date":"2026-10-29"}' \
  "$base/v1/project-code-basis")"
[[ "$response" == "422" ]] || { echo "Non-demo coordinate returned HTTP $response, expected 422" >&2; exit 1; }
readiness="$(curl --silent --show-error --max-time 5 --output /dev/null --write-out '%{http_code}' \
  "$base/v1/readiness")"
[[ "$readiness" == "404" ]] || { echo "Synthetic image exposed live readiness: HTTP $readiness" >&2; exit 1; }

# No provider bundle means startup must fail. Do not contact the HTTPS example host.
invalid="$(docker run -d --env BCM_PROVIDER_URL=https://provider.invalid \
  --env BCM_PROVIDER_AUTH_MODE=gcp_id_token \
  --env BCM_PUBLIC_ORIGIN=https://laurajoyhutchins.github.io "$image")"
for ((i=0;i<15;i++)); do
  state="$(docker inspect --format '{{.State.Running}}' "$invalid")"
  if [[ "$state" != "true" ]]; then break; fi
  sleep 1
done
state="$(docker inspect --format '{{.State.Running}}' "$invalid")"
docker rm -f "$invalid" >/dev/null 2>&1 || true
[[ "$state" != "true" ]] || { echo 'Live backend started without mandatory reviewed bundle pin' >&2; exit 1; }

echo 'BCM image built and passed synthetic-only + fail-closed startup smoke; no source-backed API deployed.'
