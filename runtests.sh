#!/usr/bin/env bash
cd /app
go mod tidy
echo "SLEEPING 5 sec..."
sleep 5 #let containers start up
echo "================================================================"
echo "Running Accounts Client Tests..."
echo "================================================================"

go test -v  ./...