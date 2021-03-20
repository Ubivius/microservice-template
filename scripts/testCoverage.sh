#!/bin/bash

echo "Running tests and displaying coverage in browser"
go test ./... -coverprofile=coverage.out -short
go tool cover -html=coverage.out
