#!/bin/bash

#Change Me!!
export WS_URL=http://service-srv.inquestdevops.com
export GROUP_ID=inquestdevops.com
export SERVICE_ID=changme
export GO_PACKAGE_PATH=com/inquestdevops/{SERVICE_ID}   # Must Match the config in /config/codegen-go-cconfig.json!!!
# shellcheck disable=SC2155
export GOPATH=$(pwd)

mkdir -p "${GOPATH}/build/tmp"
mkdir -p "${GOPATH}src/${GROUP_ID}/${SERVICE_ID}"
if [[ ! -f ${GOPATH}/build/tmp/openapi-codegen.jar ]]; then
    wget -O "${GOPATH}/build/tmp/openapi-codegen.jar" https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/6.2.1/openapi-generator-cli-6.2.1.jar
fi

java -jar "${GOPATH}/build/tmp/openapi-codegen.jar" generate\
 --input-spec ${WS_URL}/openapi.yaml\
 --generator-name go\
 --output "${GOPATH}/src/${GO_PACKAGE_PATH}" \
 --config "${GOPATH}/config/codegen-go-config.json" \
 -p enumClassPrefix=true\
 --skip-validate-spec

ClientFactoryContents="package ${SERVICE_ID}

                  import \"context\"

                  func GetClientConfiguration(authToken string, envName string) (*APIClient, context.Context) {
                  	url := \"http://localhost:8080\"
                  	if envName == \"DEV\" {
                  		url = \"http://%%{{ModuleName.lowerCase}}%%.%%{{DnsZoneDevDomain}}%%\"
                  	}
                  	if envName == \"PROD\" {
                  		url = \"%%{{ModuleName.lowerCase}}%%.%%{{DnsZoneProdDomain}}%%\"
                  	}

                  	ctx := context.WithValue(context.Background(), ContextAccessToken, authToken)

                  	resService := NewAPIClient(&Configuration{
                  		DefaultHeader: make(map[string]string),
                  		UserAgent:     \"%%{{ModuleName.lowerCase}}%%-client\",
                  		Debug:         false,
                  		Servers: ServerConfigurations{
                  			{
                  				URL:         url,
                  				Description: \"%%{{ModuleName}}%% Endpoint\",
                  			},
                  		},
                  		OperationServers: map[string]ServerConfigurations{},
                  	})
                  	return resService, ctx

                  }
"

echo "${ClientFactoryContents}" > "${GOPATH}/src/${GO_PACKAGE_PATH}/ClientFactory.java"