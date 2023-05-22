#!/usr/bin/env bash

#
# Created by zerjioang
# https://github.com/zerjioang/flights
# Copyright (c) 2021. All rights reserved.
#
# SPDX-License-Identifier: GPL-3.0
#

cd "$(dirname "$0")"

if [[ ! -f ${GOPATH}/bin/golangci-lint ]]; then
	#statements
	echo "golangci-lint not found. Downloading via curl"
	# binary will be $(go env GOPATH)/bin/golangci-lint
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
fi

${GOPATH}/bin/golangci-lint --version
# move to project root
cd ..

${GOPATH}/bin/golangci-lint run ./...