How To Use This Template
========================

##Local Developer Setup

1. Install GVM
    - ```brew install gvm``` should be enough

2. Install Go Verion >1.7.3
    - ```gvm install go1.4``` needed to build >1.4 versions
    - ```gvm use 1.4```
    - ```gvm install go1.10``` (or later)
    - ```gvm use 1.10``` (needed to run gopherjs compiler)


3. Install Delve (for Debugging)
    - Instructions [Here](https://github.com/derekparker/delve/tree/master/Documentation/installation)
    - Pay close attention to the details because this is a little sketchy.
  
    
## To Build

```
./gradlew release
```

## To Run Locally

This template can be run locally by changing to the build/gopath directory and running your main package in the bin directory.  The CWD should be the gopath root to ensure assets can be located properly.

Here is a quick example

```
cd build/gopath

./bin/server

```

The application can be viewed at ```http://localhost:8080```

If everything is working you should see text that says "Loading Application..." and get an alert saying "Hello"

## To Run Docker

This template also includes the docker facet which will automatically build a docker image for use with Cassius.  This is build automatically with the ```release``` target

```
#the <tag> is configured in the build.gradle file - it is "template-go-project:latest" by default but this should always match the module name.

/gradlew release && docker run -it -p 8080:8080 <tag>

```

## To Deploy to Cassius

The template includes a promotion properties file which should be updated with your BrightStar config details.  The realm and cell ids are required as are the different pipeline stages you which to enable by default.

This template is design to conform to the Cardinal application specs (it defines a /status endpoint and exposes a port at 8080), thus it can be promoted with the cardinal template

Promote your module by running BrightStar:

```

project module promote --configProperties promotionconfig.properties --name <module name> --majorVersion <major version> --template cardinal-app


```
