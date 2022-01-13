#!/usr/bin/env bash
cd app
go mod tidy
go test -v  ./...