#!/bin/bash

#Change Me!!
export WS_URL=http://service-srv.inquestdevops.com
export GROUP_ID=inquestdevops.com
export SERVICE_ID=changme

export GOPATH=`pwd`

mkdir -p ${GOPATH}/build/tmp
mkdir -p ${GOPATH}src/${GROUP_ID}/${SERVICE_ID}
if [[ ! -f ${GOPATH}/build/tmp/openapi-codegen.jar ]]; then
    wget -O ${GOPATH}/build/tmp/openapi-codegen.jar https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/6.2.1/openapi-generator-cli-6.2.1.jar
fi

java -jar ${GOPATH}/build/tmp/openapi-codegen.jar generate\
 --input-spec ${WS_URL}/openapi.yaml\
 --generator-name go\
 --output ${GOPATH}/src/${GROUP_ID}/${SERVICE_ID} \
 --config ${GOPATH}/config/codegen-go-config.json \
 -p enumClassPrefix=true\
 --skip-validate-spec