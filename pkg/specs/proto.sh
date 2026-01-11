#!/bin/sh

set -eu

: ${BUF:=buf}

cd "$(dirname "$0")"

if ! $BUF --version 2>/dev/null | grep -q '^[0-9]'; then
	echo "Warning: buf not found, skipping generation" >&2
	exit 0
fi

make -C ../../proto generate-go "BUF=$BUF" "OUT=$PWD"
