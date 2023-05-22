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

echo "checking source code with Go vet"

#get all files excluding vendors
filelist=$(go list ./... | grep -vendor)
for file in ${filelist}
do
	echo "static analysis of package $file"
	go vet $@ ${file}
done

echo "code checking done!"