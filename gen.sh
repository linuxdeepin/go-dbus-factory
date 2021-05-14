#!/bin/bash
set -e
export GO111MODULE=off

for dir in `find -maxdepth 1 -type d`; do
	if [ -f $dir/config.json ]; then
		./generator $dir
		pushd $dir
		echo "do "$dir
		go build -v
		popd
	fi
done
