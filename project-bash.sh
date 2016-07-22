echo "This file is intended to be sourced.  ie - 'source ./project-bash.sh'"

export GOPATH=`pwd`
export PATH=$PATH:$GOPATH/bin

echo "GOPATH set to: $GOPATH"

echo "Some Helpful Go Commands"

echo "go install github.com/<user>/hello"
echo "go get <package name>"
echo "Go Debugging: https://github.com/mailgun/godebug"
echo "godebug "

