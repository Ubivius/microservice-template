#!/bin/bash

echo "Running tests and displaying coverage in browser"
go test ./... -coverprofile=coverage.out -short
go tool cover -func=coverage.out
# -html=coverage.out to open browser with code coverage or -func=coverage.out to print results to coverage.out
