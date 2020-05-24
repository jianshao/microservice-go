#!/bin/bash

APP=$1
DIR=/home/eric/apps

if ![ -f "./go.mod"]; then
	go mod init ${APP}
fi

go build

mkdir -p ${DIR}/${APP}
cp ${APP} ${DIR}/${APP}