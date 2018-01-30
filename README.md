How To Use This Template
========================

## Local Developer Setup

This template is designed to be a starting point for GoLang development.  
There are different flavors of go projects on the different branches.  

## Dependencies

The Warbler Build plugin for go manages dependencies in two basic ways:

1. Via a Nexus Repository (for internal dependencies)
2. Third Party Dependencies managed in the Vendor folder

In the case of dependencies managed via a Maven repository - the project has a lot of control over how these dependencies
are managed, versioned, upgraded, etc.  This is the preferred way for working with code dependencies that are build internally and 
modified frequently.  Simply add these dependencies directly to the build.gradle "compile" dependencies closure
and they will be automatically fetched and exploded into your src/vendor folder by Warbler

The second option is to use standard go vendoring.  Both approaches work together.

The template recommends using the `govendor` tool for managing external dependencies.

###Installing GoVendor

```go get -u github.com/kardianos/govendor```

now ensure your GOPATH variable is set to the module root, and simply run ```govendor list```

You can add any additional dependencies via ```govendor fetch <dep>```

## To Generate Client

There is a shell script `refresh-api.sh` which needs some service specific modifications.

Update the values for the $GROUP_ID, $SERVICE_ID and $WS_URL variables.  Then run `./refresh-api`.  The script will
download the generator automatically from Maven and generate the client SDK.

##To Build

1. Install GVM
    - ```brew install gvm``` should be enough

2. Install Go Verion >1.7.3
    - ```gvm install go1.4``` needed to build >1.4 versions
    - ```gvm use 1.4```
    - ```gvm install go1.11.5``` (or later)


3. Install Delve (for Debugging)
    - Instructions [Here](https://github.com/derekparker/delve/tree/master/Documentation/installation)
    - Pay close attention to the details because this is a little sketchy.
  
## To Setup Swagger Settings
This template assumes that the service you are generating an SDK for is running locally at port 8080.  This can be configured.  The config codegen file will specify the default base url of the running production service whereas the refresh-api.sh file requires the URL of the service you wish to use to generate the SDK. 
1. Modify ```./config/codegen-go-config.json``` file to specify the correct service url and package settings.

2. Review ```refresh-api.sh``` file settings to be sure they match your service.  
    
## To Build

```
./gradlew release
```
