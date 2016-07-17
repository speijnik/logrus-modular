#!/bin/sh

THIS_PATH=$(dirname $0)
cd ${THIS_PATH}
go test -v -cover -coverprofile coverage.out ./ || exit 1
go tool cover -html coverage.out -o coverage.html || exit 1
exit 0
