#!/bin/bash

if [ "$1" != "golua" ] && [ "$1" != "ollama" ]; then
    echo "Usage: $0 [golua|ollama] <GITHUB_TOKEN> <VERSION>"
    echo "  golua   - Build golua runtime image"
    echo "  ollama  - Build ollama runtime image"
    echo "  VERSION - Image version tag (e.g., v1.0.0, latest)"
    exit 1
fi

if [ -z "$2" ]; then
    echo "No GITHUB_TOKEN"
    echo "Usage: $0 [golua|ollama] <GITHUB_TOKEN> <VERSION>"
    exit 1
fi

if [ -z "$3" ]; then
    echo "No VERSION specified"
    echo "Usage: $0 [golua|ollama] <GITHUB_TOKEN> <VERSION>"
    exit 1
fi

if [ "$1" = "golua" ]; then
    docker build --progress=plain \
        --build-arg GITHUB_TOKEN="$2" \
        -f Dockerfile.golua \
        -t chriswebber/docker-golua:"$3" .
else
    docker build --progress=plain \
        --build-arg GITHUB_TOKEN="$2" \
        -f Dockerfile.ollama \
        -t chriswebber/docker-ollama:"$3" .
fi