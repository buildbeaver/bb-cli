#!/bin/bash
set -e
if [ -n "${BB_DEBUG}" ]; then
  set -x
fi

ROOT_DIR=$(realpath "$(git rev-parse --show-toplevel)")
. "${ROOT_DIR}/build/scripts/lib/go-env.sh"
check_deps "go"

# Specify our ldflags for injecting our version information into our binaries.
PKG="github.com/buildbeaver/bb"
VERSION_INFO=$(${ROOT_DIR}/build/scripts/version-info.sh)
GIT_SHA_SHORT=$(${ROOT_DIR}/build/scripts/version-info.sh sha-short)
VERSION_VAR="-X '${PKG}/common/version.VERSION=${VERSION_INFO}' -X '${PKG}/common/version.GITCOMMIT=${GIT_SHA_SHORT}'"
GO_LDFLAGS="-ldflags=${VERSION_VAR}"

for cmd_dir in backend/*/cmd/*; do
  bin_name="$(basename "${cmd_dir}")"
  bin_out="${GOBIN}/${bin_name}"
  pushd "${cmd_dir}"
    echo "Building: ${bin_name} > ${bin_out}"
    go build "${GO_LDFLAGS}" -mod=vendor -o "${bin_out}" .
  popd
done