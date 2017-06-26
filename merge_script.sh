#!/usr/bin/env bash

PULL_REQUEST="$1"
echo ${PULL_REQUEST}
curl -X PUT -H "Authorization: Basic ZGFuaWVsc3Vzc2FAZ21haWwuY29tOkVkdXRlYW1vWzExXQ==" https://api.github.com/repos/danielsussa/simpleCRUD/pulls/${PULL_REQUEST}/merge
