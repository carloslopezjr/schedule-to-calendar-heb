#!/bin/bash


echo "Sending schedule to Google Calendar..."

# Navigate to the script's directory
cd "$(dirname "$0")" || exit

# run go files
go run main.go parse.go calendar.go

