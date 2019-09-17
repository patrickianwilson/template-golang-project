package properties

//Config Property Keys
const (
	GIT_EXECUTABLE                   = "sys.git.path"
	GITHUB_ACCESS_TOKEN              = "provider.repos.github.token"
	GITHUB_ORGANIZATION              = "provider.repos.github.organization"
	MODULE_SERVICE_URL               = "provider.internal.service.module.url"          // this is the customer module service url
	MODULE_SERVICE_UPSTREAM_URL      = "provider.internal.service.module.upstream.url" //this is the url for the "upstream" module service (IE = Inquest Devops)
	CONVENTIONS_SERVICE_URL          = "provider.internal.service.conventions.url"
	DEFAULT_GROUP_ID_KEY             = "defaults.group.id"
	MODULE_SERVICE_PASSWORD          = "provider.internal.service.module.password"
	MODULE_SERVICE_UPSTREAM_PASSWORD = "provider.internal.service.module.upstream.password"
)

//Defaults
const (
	NO_MODULE_SERVICE                   = "module-service-url-not-set"
	NO_UPSTREAM_MODULE_SERVICE          = "upstream-module-service-url-not-set"
	NO_MODULE_SERVICE_PASSWORD          = "module-service-password-not-set"
	NO_UPSTREAM_MODULE_SERVICE_PASSWORD = "upstream-module-service-password-not-set"
	NO_BRIGHTSTAR_SERVICE               = "brightstar-service-url-not-set"
	NO_DEFAULT_GROUP_ID                 = "default-group-id-not-set"
)
