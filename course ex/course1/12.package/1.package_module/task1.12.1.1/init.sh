#!/bin/bash


if [ -n "$1" ]
then
	go mod init $1
else
	echo "Module name argument is missing"
	exit 1
fi
