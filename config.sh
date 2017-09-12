
#!/bin/bash
#
# Setup for GOPATH variable
# Import all frameworks and packages from github
# build and run the gocouch database
#
# Author - Dragos STOICA
# Date: 23.02.2016
#

export GOPATH=`pwd`
export GOBIN=$GOPATH/bin
mkdir -p bin
mkdir -p pkg

#git clone https://github.com/boltdb/bolt.git ./src/github.com/boltdb/bolt
#git clone https://github.com/labstack/echo.git ./src/github.com/labstack/echo
#git clone https://github.com/labstack/gommon.git ./src/github.com/labstack/gommon
#git clone https://github.com/mattn/go-colorable.git ./src/github.com/mattn/go-colorable
#git clone https://github.com/mattn/go-isatty.git ./src/github.com/mattn/go-isatty
#git clone https://github.com/golang/net.git ./src/golang.org/x/net
#git clone https://github.com/golang/crypto.git ./src/golang.org/x/crypto
#git clone https://github.com/golang/text.git ./src/golang.org/x/text

go get -u github.com/coreos/bbolt/...
go get -u github.com/labstack/echo/...
go get -u golang.org/x/crypto/...
go get -u golang.org/x/net ...
go get -u golang.org/x/text ...

go build
./gocouch
