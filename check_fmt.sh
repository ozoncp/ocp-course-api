#!/bin/bash

bad=$(gofmt -l -e .)
if [ $bad ]; then
	echo $bad
	echo fmt - FAIL
	exit 1
fi

echo fmt - OK
