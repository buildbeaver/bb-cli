#!/bin/bash
set -e
if [ -n "${BB_DEBUG}" ]; then
  set -x
fi

ROOT_DIR=$(realpath "$(git rev-parse --show-toplevel)")
. "${ROOT_DIR}/build/scripts/lib/env.sh"

export GODIR="${BUILD_DIR}/go"
export GOBIN="${GODIR}/bin"
export GOCACHE="${GODIR}/cache"

mkdir -p "${GOCACHE}" "${GOBIN}"
