#!/bin/bash

if [[ $# -eq 0 ]]; then
    echo "usage: ./build.sh {simples|example1}"
    exit 1
fi

if [[ $1 = "simple" ]]; then
    cargo run --bin simple
    exit 0
fi

if [[ $1 = "example1" ]]; then
    cargo run --bin example1
    exit 0
fi