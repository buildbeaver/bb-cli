#!/bin/bash
set -e
if [ -n "${BB_DEBUG}" ]; then
  set -x
fi

ROOT_DIR=$(realpath "$(git rev-parse --show-toplevel)")
. "${ROOT_DIR}/build/scripts/lib/env.sh"

NODE_PATH="${BUILD_DIR}/node/node_modules"
export NODE_PATH
export PATH="${PATH}:${NODE_PATH}/.bin"