#!/bin/bash
set -e
if [ -n "${BB_DEBUG}" ]; then
  set -x
fi

ROOT_DIR=$(realpath "$(git rev-parse --show-toplevel)")
. "${ROOT_DIR}/build/scripts/lib/env.sh"
check_deps "openapi-generator"

openapi-generator generate \
  -i "backend/server/api/rest/openapi/$1" \
  -g "$2" \
  --additional-properties disallowAdditionalPropertiesIfNotPresent=false \
  --additional-properties packageName=client \
  --additional-properties isGoSubmodule=true \
  --additional-properties modelPropertyNaming=snake_case \
  -o "$3"

if [ "$2" = "go" ]; then
  rm "$3"/go.{mod,sum}
  rm -rf "$3/test"
fi
