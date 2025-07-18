#!/bin/bash

set -e

cd "$(dirname "$0")"

if [[ "$1" == "--reset" ]]; then
    echo "Reset flag detected. Removing lock file and reinitializing..."
    go run main.go --reset
else
    go run main.go
fi
