#!/bin/bash

if ! command -v go &> /dev/null
then
    echo "Go is not installed. Please install Go first."
    exit 1
fi

go install github.com/tuhinexe/optical@latest


if [ $? -eq 0 ]; then
    echo "Optical CLI tool installed successfully!"
    echo "You can now use 'optical -create' to create a new Optical project."
else
    echo "Installation failed. Please check your Go environment and try again."
    exit 1
fi


if ! command -v optical &> /dev/null
then
    echo "Warning: The installation directory might not be in your PATH."
    echo "Add the Go bin directory to your PATH to use optical from anywhere."
fi