#!/usr/bin/env bash
cd /app
go mod tidy
echo "SLEEPING 5 sec..."
sleep 5 #let containers start up
echo "================================================================"
echo "Running Accounts Client Tests..."
echo "================================================================"


if go test -v ./... ; then
	echo "================================================================"
	echo "ALL TESTS PASSED..."
	echo "================================================================"

else
	echo "====================================================================================================="
	echo "TESTS FAILED! Please verify correct API_HOST enviroment variable is configured in docker-compose.yml "
	echo "====================================================================================================="
fi