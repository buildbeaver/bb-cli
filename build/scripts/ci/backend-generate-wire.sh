#!/bin/bash
set -e
if [ -n "${BB_DEBUG}" ]; then
  set -x
fi

ROOT_DIR=$(realpath "$(git rev-parse --show-toplevel)")
. "${ROOT_DIR}/build/scripts/lib/go-env.sh"
check_deps "wire.exe"

for wire_file in backend/*/app/wire.go; do
  pushd "$(dirname "${wire_file}")"
    wire.exe
  popd
done