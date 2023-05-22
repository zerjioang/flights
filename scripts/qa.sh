#!/usr/bin/env bash

#
# Created by zerjioang
# https://github.com/zerjioang/flights
# Copyright (c) 2021. All rights reserved.
#
# SPDX-License-Identifier: GPL-3.0
#

cd "$(dirname "$0")"

bash fmt.sh
bash simplify.sh
bash goimports.sh

bash govet.sh
bash golangci.sh