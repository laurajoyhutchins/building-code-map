#!/bin/sh
set -eu

package_root=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)
cd "$package_root"

test -x bin/bcm
test -f manifests/bundle.json
test -f manifests/checksums.json

checksum_lines=$(sed -n 's/^    "\([^"]*\)": "\([0-9a-f][0-9a-f]*\)",\{0,1\}$/\2  \1/p' manifests/checksums.json)
[ -n "$checksum_lines" ] || { echo "no checksum entries found" >&2; exit 4; }

printf '%s\n' "$checksum_lines" | while IFS='  ' read -r digest path; do
  [ "${#digest}" -eq 64 ] || { echo "invalid checksum digest: $path" >&2; exit 4; }
  [ -f "$path" ] || { echo "missing checksummed file: $path" >&2; exit 4; }
  actual=$(sha256sum "$path" | awk '{print $1}')
  [ "$actual" = "$digest" ] || { echo "checksum mismatch: $path" >&2; exit 4; }
done

bin/bcm inspect bundle --bundle manifests/bundle.json >/dev/null
bin/bcm resolve --bundle manifests/bundle.json --point=-104.99,39.74 --as-of=2026-08-06 >/dev/null
printf '%s\n' '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{}}' | bin/bcm serve --bundle manifests/bundle.json --mcp-stdio >/dev/null
echo "offline verification passed"
