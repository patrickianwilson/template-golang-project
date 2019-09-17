package properties

import (
	"encoding/json"
	"log"
	"os"
)

type ConfigProperties interface {
	GetConfigProperty(key string) string
	OverlayConfigFromFile(cf string)
}

type HardCodedConfigProperties struct {
	configMap map[string]string
}

func (p *HardCodedConfigProperties) GetConfigProperty(key string) string {
	return p.configMap[key]
}

/*
Eventually, this should be loaded from a properties file in ~/.brightstar/config
*/
func NewHardCodedConfigProperties() ConfigProperties {

	result := HardCodedConfigProperties{
		configMap: map[string]string{
			GIT_EXECUTABLE:                   "/usr/bin/git",
			GITHUB_ACCESS_TOKEN:              "access-token-not-set",
			GITHUB_ORGANIZATION:              "",
			MODULE_SERVICE_URL:               NO_MODULE_SERVICE,
			DEFAULT_GROUP_ID_KEY:             NO_DEFAULT_GROUP_ID,
			CONVENTIONS_SERVICE_URL:          NO_BRIGHTSTAR_SERVICE,
			MODULE_SERVICE_PASSWORD:          NO_MODULE_SERVICE_PASSWORD,
			MODULE_SERVICE_UPSTREAM_URL:      NO_UPSTREAM_MODULE_SERVICE,
			MODULE_SERVICE_UPSTREAM_PASSWORD: NO_UPSTREAM_MODULE_SERVICE_PASSWORD,
		},
	}
	return &result
}

//This is the format for the config file.  Not all config properties
//will be supported so we keep a different struct for these to make
//JSON serialization easier.
type ConfigPropertyFileStruct struct {
	GithubToken                   string
	ModuleServiceUrl              string
	BrightstarServiceUrl          string
	GithubOrganization            string
	DefaultGroupId                string
	GitExecutable                 string
	ModuleServicePassword         string
	UpstreamModuleServicePassword string
	UpstreamModuleServiceUrl      string
}

func (p *HardCodedConfigProperties) OverlayConfigFromFile(cf string) {
	file, err := os.Open(cf)
	defer file.Close()

	if err != nil {
		log.Printf("Warning!  Unable to open file %s, error %s", cf, err)
		log.Printf("Continuing with default configuration")
		return
	}

	config := &ConfigPropertyFileStruct{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)

	if err != nil {
		log.Printf("Unable to parse config json content: %s", err)
	}

	if config.GithubToken != "" {
		p.configMap[GITHUB_ACCESS_TOKEN] = config.GithubToken
	}
	if config.BrightstarServiceUrl != "" {
		p.configMap[CONVENTIONS_SERVICE_URL] = config.BrightstarServiceUrl
	}
	if config.ModuleServiceUrl != "" {
		p.configMap[MODULE_SERVICE_URL] = config.ModuleServiceUrl
	}
	if config.DefaultGroupId != "" {
		p.configMap[DEFAULT_GROUP_ID_KEY] = config.DefaultGroupId
	}
	if config.GitExecutable != "" {
		p.configMap[GIT_EXECUTABLE] = config.GitExecutable
	}
	if config.GithubOrganization != "" {
		p.configMap[GITHUB_ORGANIZATION] = config.GithubOrganization
	}
	if config.ModuleServicePassword != "" {
		p.configMap[MODULE_SERVICE_PASSWORD] = config.ModuleServicePassword
	}
	if config.UpstreamModuleServiceUrl != "" {
		p.configMap[MODULE_SERVICE_UPSTREAM_URL] = config.UpstreamModuleServiceUrl
	}
	if config.UpstreamModuleServicePassword != "" {
		p.configMap[MODULE_SERVICE_UPSTREAM_PASSWORD] = config.UpstreamModuleServicePassword
	}
}
