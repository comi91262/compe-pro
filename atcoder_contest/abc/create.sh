#! /usr/bin/env bash

if [ $# -ne 1 ]; then
    echo "error"
    exit 1
fi

mkdir $1
touch $1/{a,b,c,d,e,f}.go
