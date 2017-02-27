#!/bin/bash

go test -v -coverprofile=coverage.out ./cwb/
go tool cover -html=coverage.out -o coverage.html
open coverage.html