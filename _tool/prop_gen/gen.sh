#!/bin/sh
set -x
export GO111MODULE=off
dstFile=$proj_deepin_go_lib/dbusutil/proxy/prop.go
go run main.go > $dstFile
goimports -w $dstFile
