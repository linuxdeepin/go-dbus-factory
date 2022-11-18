#!/bin/bash
set -e

for dir in `find -maxdepth 2 -type d`; do
	if [ -f $dir/config.json ]; then
		./generator $dir
		pushd $dir
		go build -v
		popd
	fi
done
