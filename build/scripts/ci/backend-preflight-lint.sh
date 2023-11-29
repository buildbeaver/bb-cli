#!/bin/bash
set -e
if [ -n "${BB_DEBUG}" ]; then
  set -x
fi

ROOT_DIR=$(realpath "$(git rev-parse --show-toplevel)")
. "${ROOT_DIR}/build/scripts/lib/node-env.sh"
check_deps "goimports.exe"

cd "backend"
out="$(find . -type f -name '*.go' -not -path '*/vendor/*' -not -path '*/wire_gen.go' -exec goimports.exe -d {} \;)"
if [ "${out}" != "" ]; then
  echo ""
  echo "Looks like you forgot to run 'goimports' before committing the following files:"
  echo "${out}"
  exit 1
fi
echo "No linting issues found"