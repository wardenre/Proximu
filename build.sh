#!/bin/bash

if ! command -v go &> /dev/null; then
    echo "Go is not installed. Please install Go: https://go.dev/dl/"
    exit 1
fi

mkdir -p build

echo "Building Go Proximu..."
go build -o build/Proximu

if [ $? -eq 0 ]; then
    echo "Build successful: build/Proximu"
else
    echo "Build failed."
    exit 1
fi