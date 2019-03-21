How To Use This Template
========================

##Local Developer Setup

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

##To Build

1. Install GVM
    - ```brew install gvm``` should be enough

2. Install Go Verion >1.7.3
    - ```gvm install go1.4``` needed to build >1.4 versions
    - ```gvm use 1.4```
    - ```gvm install go1.11.5``` (or later)
    
## To Build

```
./gradlew release
```
