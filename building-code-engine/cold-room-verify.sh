#!/bin/sh
set -eu

source_root=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)
temporary_root=$(mktemp -d)
trap 'rm -rf "$temporary_root"' EXIT
cp -R "$source_root" "$temporary_root/package"

network_isolator="none"
if command -v unshare >/dev/null 2>&1 \
  && command -v ip >/dev/null 2>&1 \
  && unshare -n sh -c 'ip link set lo up' >/dev/null 2>&1; then
  network_mode="network-namespace-disabled"
  network_isolator="network"
elif command -v unshare >/dev/null 2>&1 \
  && command -v ip >/dev/null 2>&1 \
  && unshare -Urn sh -c 'ip link set lo up' >/dev/null 2>&1; then
  network_mode="network-namespace-disabled"
  network_isolator="user-network"
else
  network_mode="network-namespace-fallback"
fi

printf '%s\n' "$network_mode"
printf 'network-namespace-isolator=%s\n' "$network_isolator"
cd "$temporary_root/package"

run() {
  case "$network_isolator" in
    network)
      unshare -n sh -c 'ip link set lo up; exec env -i PATH=/usr/bin:/bin HOME=/tmp "$@"' sh "$@"
      ;;
    user-network)
      unshare -Urn sh -c 'ip link set lo up; exec env -i PATH=/usr/bin:/bin HOME=/tmp "$@"' sh "$@"
      ;;
    *)
      env -i PATH=/usr/bin:/bin HOME=/tmp "$@"
      ;;
  esac
}

probe_json() {
  label=$1
  shift
  output=$(mktemp)
  set +e
  run "$@" >"$output"
  status=$?
  set -e
  if [ "$status" -ne 0 ]; then
    echo "$label failed with exit code $status" >&2
    cat "$output" >&2
    rm -f "$output"
    exit "$status"
  fi
  run python3 -c 'import json, sys; json.load(open(sys.argv[1], encoding="utf-8"))' "$output"
  rm -f "$output"
}

run ./verify-offline.sh
probe_json bundle-inspect ./bin/bcm inspect bundle --bundle manifests/bundle.json
probe_json point-resolve ./bin/bcm resolve \
  --bundle manifests/bundle.json \
  --point=-104.99,39.74 \
  --as-of=2026-08-06
probe_json geocode ./bin/bcm geocode \
  --bundle manifests/bundle.json \
  --address '1600 N Broadway St, Denver, CO 80202'
probe_json address-resolve ./bin/bcm resolve \
  --bundle manifests/bundle.json \
  --address '1600 N Broadway St, Denver, CO 80202' \
  --as-of=2026-08-06

port=18765
run sh -c '
  set -eu
  port=$1
  ./bin/bcm serve --bundle manifests/bundle.json --http "127.0.0.1:$port" \
    >bcm-cold-room-http.log 2>bcm-cold-room-http.err &
  server_pid=$!
  trap '\''kill "$server_pid" 2>/dev/null || true'\'' EXIT

  i=0
  while ! curl --noproxy "*" -fsS "http://127.0.0.1:$port/v1/readiness" >/dev/null 2>&1; do
    i=$((i + 1))
    [ "$i" -lt 50 ] || { cat bcm-cold-room-http.err >&2; exit 1; }
    sleep 0.1
  done

  curl --noproxy "*" -fsS "http://127.0.0.1:$port/v1/resolve" \
    -H "Content-Type: application/json" \
    --data '\''{"point":{"longitude":-104.99,"latitude":39.74},"applicability_date":"2026-08-06"}'\'' \
    >/dev/null
' sh "$port"

printf '%s\n' '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{}}' | run ./bin/bcm serve --bundle manifests/bundle.json --mcp-stdio >/dev/null
echo "cold-room verification passed"
