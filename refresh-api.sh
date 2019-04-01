#!/bin/bash

export WS_URL=http://modules.wilsonsinquest.com
export GROUP_ID=mygroup
export SERVICE_ID=myservice

export GOPATH=`pwd`

mkdir -p ${GOPATH}/build/tmp
mkdir -p ${GOPATH}src/clientsdk/moduleservice
if [[ ! -f ${GOPATH}/build/tmp/openapi-codegen.jar ]]; then
    wget -O ${GOPATH}/build/tmp/openapi-codegen.jar http://central.maven.org/maven2/org/openapitools/openapi-generator-cli/3.3.4/openapi-generator-cli-3.3.4.jar
fi

java -jar ${GOPATH}/build/tmp/openapi-codegen.jar generate\
 --input-spec ${WS_URL}/openapi.yaml\
 --generator-name go\
 --output ${GOPATH}/src/github.com/$GROUP_ID/$SERVICE_ID \
 --config ${GOPATH}/config/codegen-go-config.json
