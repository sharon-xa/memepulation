#!/bin/bash

mkdir -p builds

# List of platforms and architectures to build for
platforms=("linux" "darwin")
architectures=("amd64" "arm64")

for platform in "${platforms[@]}"; do
  for arch in "${architectures[@]}"; do
    output_name="builds/memepulation-$platform-$arch"
    
    echo "Building for $platform/$arch..."

    env GOOS=$platform GOARCH=$arch go build -o $output_name

    if [ $? -ne 0 ]; then
      echo "Failed to build for $platform/$arch"
      exit 1
    fi

    echo "Build successful: $output_name"
  done
done

echo "All builds completed and stored in the 'builds' directory."
