#!/bin/bash

#
# Created by zerjioang
# https://github.com/zerjioang/flights
# Copyright (c) 2021. All rights reserved.
#
# SPDX-License-Identifier: GPL-3.0
#

cd "$(dirname "$0")"

# move to project root dir from ./scripts to ./
cd ..

echo "checking imports in source code with goimports"

echo "checking if goimports is installed in $GOPATH"
if [[ ! -f ${GOPATH}/bin/goimports ]]; then
	#statements
	echo "goimports not found. Downloading via go get"
	go install golang.org/x/tools/cmd/goimports@latest
fi

#get all files excluding vendors
filelist=$(find . -type f -name "*.go" | grep -vendor)
for file in ${filelist}
do
	echo "goimports check in file $file"
	${GOPATH}/bin/goimports -v -w ${file}
done

echo "code formatting done!"