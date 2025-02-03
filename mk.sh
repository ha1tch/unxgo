#!/bin/bash
mkdir -p bin

echo "Building all binaries for zxa..."
pushd cmd/uxn

BASENAME="unx"

# Build for Windows
GOOS=windows GOARCH=amd64 go build -x -o ../../bin/$BASENAME.win64.exe   main.go
GOOS=windows GOARCH=386   go build -x -o ../../bin/$BASENAME.win32.exe   main.go

# Build for Linux
GOOS=linux   GOARCH=amd64 go build -x -o ../../bin/$BASENAME.linux64     main.go
GOOS=linux   GOARCH=386   go build -x -o ../../bin/$BASENAME.linux32     main.go

# Build for macOS (modern architectures)
GOOS=darwin  GOARCH=arm64 go build -x -o ../../bin/$BASENAME.mac64.m1    main.go
GOOS=darwin  GOARCH=amd64 go build -x -o ../../bin/$BASENAME.mac64.intel main.go

# Build for Raspberry Pi
GOOS=linux   GOARCH=arm   GOARM=6  go build -x -o ../../bin/$BASENAME.rpi.arm6   main.go  # Pi 1, Pi Zero
GOOS=linux   GOARCH=arm   GOARM=7  go build -x -o ../../bin/$BASENAME.rpi.arm7   main.go  # Pi 2, Pi 3 (32-bit)
GOOS=linux   GOARCH=arm64          go build -x -o ../../bin/$BASENAME.rpi.arm64  main.go  # Pi 3, Pi 4, Pi 5 (64-bit)

# ---------------------------------------------------------------
# Important Note on 32-bit macOS (i386) Builds
# ---------------------------------------------------------------
# Go 1.15 was the last version to support building 32-bit (i386) binaries for macOS.
# If you need to generate 32-bit binaries for older Intel Macs (2006–2010 era),
# you must use Go 1.15 or an earlier version.
#
# Starting from Go 1.16, support for 32-bit macOS was officially removed.
# Newer Go versions (1.16+) only compile 64-bit binaries for macOS.
#
# Reference: https://golang.org/doc/go1.16#darwin
#
# If you STILL need to generate 32-bit Intel binaries for macOS and have installed
# Go 1.15 (or earlier), you can attempt the following build command:
#
# GOOS=darwin  GOARCH=386 go build -x -o ../../bin/$BASENAME.mac32.intel main.go
#
# However, this is completely unsupported and untested in modern Go versions.
# There are no guarantees that this will work.
# ---------------------------------------------------------------

popd
