#!/usr/bin/env bash

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

echo "Checking source files license header..."
pwd

copyrightContent=$(cat ./embed/LICENSE.txt)
echo "default header copyright content is:"
echo $copyrightContent

files=$(find . -type f -name "*.go" | grep -vendor)
for f in $files
do
	echo "checking license status of $f"
	if ! grep -q 'Copyright' $f; then
		cat ./embed/LICENSE.txt $f > $f.new && mv $f.new $f
	fi
done