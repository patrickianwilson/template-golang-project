#!/bin/bash
./project-developer-setup.sh

export GOPATH=`pwd`
go get github.com/gorilla/mux
go get github.com/mailgun/godebug
