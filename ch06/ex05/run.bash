#!/bin/bash

export GOARCH=amd64
echo "testing with" $GOARCH
go test

export GOARCH=386
echo "testing with" $GOARCH
go test

export GOARCH=amd64
