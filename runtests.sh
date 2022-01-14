#!/usr/bin/env bash
cd /app
go mod tidy
echo "SLEEPING 5 sec..."
sleep 5 #let containers start up
echo "================================================================"
echo "Running Accounts Client Tests..."
echo "================================================================"


if go test -v ./... ; then
	echo "ALL TESTS PASSED..."
else
	echo "====================================================================================================="
	echo "TESTS FAILED! Please verify correct API_HOST enviroment variable is configured in docker-compose.yml "
	echo "====================================================================================================="
fi