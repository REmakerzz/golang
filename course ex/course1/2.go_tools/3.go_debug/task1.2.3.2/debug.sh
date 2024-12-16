#!/bin/bash
go build -o myprogram main.go 
echo "Debug started..."
dlv exec ./myprogram
echo "Debug ended."
