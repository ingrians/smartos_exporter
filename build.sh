#!/bin/bash
mkdir $GOPATH/bin/solaris
GOOS=solaris
export GOOS
go get golang.org/x/sys/unix
GOARCH=amd64 go build -o $GOPATH/bin/solaris/smartos_exporter github.com/ingrians/smartos_exporter
