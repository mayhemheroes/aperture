#!/usr/bin/env bash
set -eux

# This script builds a Go binary and injects build-time variables.
# TODO: it should be used in every Dockerfile requiring version endpoint.

VERSION=${VERSION:-0.0.1}
BUILD_TIME=$(date --rfc-3339=seconds)
GOOS=$(go env GOOS)
GOARCH=$(go env GOARCH)
HOSTNAME=$(hostname)
SERVICE_FILE=$(basename -- "${TARGET}")
SERVICE="${SERVICE_FILE%.*}"

LDFLAGS="\
    ${LDFLAGS:-} \
    -X 'github.com/fluxninja/aperture/pkg/info.Version=${VERSION}' \
    -X 'github.com/fluxninja/aperture/pkg/info.Service=${SERVICE}' \
    -X 'github.com/fluxninja/aperture/pkg/info.BuildHost=${HOSTNAME}' \
    -X 'github.com/fluxninja/aperture/pkg/info.BuildOS=${GOOS}/${GOARCH}' \
    -X 'github.com/fluxninja/aperture/pkg/info.BuildTime=${BUILD_TIME}' \
    -X 'github.com/fluxninja/aperture/pkg/info.GitBranch=${GIT_BRANCH}' \
    -X 'github.com/fluxninja/aperture/pkg/info.GitCommitHash=${GIT_COMMIT_HASH}' \
    -X 'github.com/fluxninja/aperture/pkg/info.Prefix=${PREFIX}' \
"
if [ -z "${RACE:-}" ]; then
    go build --ldflags "${LDFLAGS}" -o "${TARGET}" "${SOURCE}"
else
    go build --race --ldflags "${LDFLAGS}" -o "${TARGET}" "${SOURCE}"
fi
