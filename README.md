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
   

4. Set your ```GOPATH``` variable to the /Service directory.
    - ```cd ./Service```
    - ```export GOPATH=$GOPATH:`pwd` ```
    - ```export PATH=$PATH:$GOPATH/bin``` needed to get tools fetched via ```go get``` onto your path
    
5. Build and Run the project with 
    - ```go build main && ./main```
    
