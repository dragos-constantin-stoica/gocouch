
#!/bin/bash
#
# Setup for GOPATH and GOBIN environment variables
# Import all frameworks and packages from github and golang
# build and run the gocouch database
#
# Author - Dragos STOICA
# Date: 23.02.2016
#

# Clean and exit - this must be performed before commit
if [ $# -eq 1 ] && [ "$1" == "clean" ]; then
    echo "Cleanup the directories and executable"
	rm -fr bin pkg src/github.com src/golang.org gocouch
	exit 0
fi

# Else get the depencies and build the executable
export GOPATH=`pwd`
export GOBIN=$GOPATH/bin
# Cleanup the folders and executable
rm -fr bin pkg src/github.com src/golang.org gocouch
# Create necessay structure
mkdir -p bin
mkdir -p pkg

# Keep BoltDB stable final version. There is an alternative proposed on the github page
go get -u github.com/boltdb/bolt/...
go get -u github.com/labstack/echo/...
go get -u golang.org/x/crypto/...
go get -u golang.org/x/net ...
go get -u golang.org/x/text ...

go build
./gocouch
