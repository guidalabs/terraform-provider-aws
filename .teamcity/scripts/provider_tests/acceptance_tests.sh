#!/usr/bin/env bash

# Code generated by internal/generate/teamcity/provider_tests.go; DO NOT EDIT.

set -euo pipefail

# shellcheck disable=2157 # This isn't a constant string, it's a TeamCity variable substitution
if [[ -n "%ACCTEST_ROLE_ARN%" ]]; then
    conf=$(pwd)/aws.conf

    function cleanup {
        rm "${conf}"
    }
    trap cleanup EXIT

    touch "${conf}"
    chmod 600 "${conf}"
    cat <<EOF >"${conf}"
[profile primary]
role_arn       = %ACCTEST_ROLE_ARN%
source_profile = primary_user

[profile primary_user]
aws_access_key_id     = %AWS_ACCESS_KEY_ID%
aws_secret_access_key = %AWS_SECRET_ACCESS_KEY%
EOF

    unset AWS_ACCESS_KEY_ID
    unset AWS_SECRET_ACCESS_KEY

    export AWS_CONFIG_FILE="${conf}"
    export AWS_PROFILE=primary
fi

TF_ACC=1 go test \
    ./internal/acctest/... \
    ./internal/attrmap/... \
    ./internal/conns/... \
    ./internal/create/... \
    ./internal/dns/... \
    ./internal/enum/... \
    ./internal/envvar/... \
    ./internal/errs/... \
    ./internal/experimental/... \
    ./internal/flex/... \
    ./internal/framework/... \
    ./internal/function/... \
    ./internal/generate/... \
    ./internal/io/... \
    ./internal/json/... \
    ./internal/logging/... \
    ./internal/maps/... \
    ./internal/namevaluesfilters/... \
    ./internal/provider/... \
    ./internal/reflect/... \
    ./internal/retry/... \
    ./internal/sdkv2/... \
    ./internal/semver/... \
    ./internal/slices/... \
    ./internal/sweep/... \
    ./internal/tags/... \
    ./internal/tfresource/... \
    ./internal/types/... \
    ./internal/vault/... \
    ./internal/verify/... \
    ./internal/yaml/... \
    -json -count=1 -parallel "%ACCTEST_PARALLELISM%" -timeout=0 -run=TestAcc
