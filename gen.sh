#!/bin/bash

for dir in `find -maxdepth 1 -type d`; do
	if [ -f $dir/config.json ]; then
		./generator $dir
	fi
done
