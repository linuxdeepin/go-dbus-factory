#!/bin/sh
set -x
dstFile=$proj_deepin_go_lib/dbusutil/proxy/prop.go
go run main.go > $dstFile
go fmt $dstFile
