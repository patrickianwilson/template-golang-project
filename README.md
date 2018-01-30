How To Use This Template
========================

##Local Developer Setup

1. Install GVM
    - ```brew install gvm``` should be enough

2. Install Go Verion >1.7.3
    - ```gvm install go1.4``` needed to build >1.4 versions
    - ```gvm use 1.4```
    - ```gvm install go1.7.3``` (or later)


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