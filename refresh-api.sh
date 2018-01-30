#!/bin/bash

export WS_URL=http://localhost:8080

export GOPATH=`pwd`

mkdir -p ${GOPATH}/build/tmp

if [ ! -f ${GOPATH}/build/tmp/codegen.jar ]; then
    wget -O ${GOPATH}/build/tmp/codegen.jar http://central.maven.org/maven2/io/swagger/swagger-codegen-cli/2.2.3/swagger-codegen-cli-2.2.3.jar
fi

java -jar ${GOPATH}/build/tmp/codegen.jar generate -i ${WS_URL}/swagger.yaml -l go -o ${GOPATH}/src/github.com/patrickianwilson/cassiusservice -c ${GOPATH}/config/codegen-go-config.json

