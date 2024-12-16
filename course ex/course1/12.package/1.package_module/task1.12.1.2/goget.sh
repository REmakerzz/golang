#!/bin/bash


if [ -n "$1" ]
then
	go mod init $1
	go get github.com/yuin/goldmark
else
	echo "Необходимо указать имя модуля"
	exit
fi
