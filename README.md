# gocouch
CouchDB implementation in Go. The main target platform is for mobile devices.  
The name is not final, it represents a concatenation of Go and Couch. We are still thinking to call it Gr0_0uchO DB, and add a mustache on a CouchDB logo.  
This project is not another implementation of CouchDB system. It came as an alternative to existing implementations: PouchDB, CouchBase Lite, TouchDB etc for mobile devices.
The main purpose is to allow a decent storage mechanism of data on mobile devices that will be able to replicate/sync with a CouchDB sever. 
By its nature GocouchDB runs off-line, stores data in JSON format and exposes a REST API to the application, and most important it syncs.

# 0. Introduction
Get this repository from github: <code>git clone https://github.com/iqcouch/gocouch.git .</code>
Setup environment variables:  
<code>cd gocouch/gocouch; export GOPATH=\`pwd\`</code>  
The storage is based on BoltDB (https://github.com/boltdb/bolt). You need to add the project from github  
<code> 
git clone https://github.com/boltdb/bolt.git ./src/github.com/boltdb/bolt
</code>  
The REST API is assured by echo framework (http://labstack.com/echo and https://github.com/labstack/echo). Add this framework from github <code>git clone https://github.com/labstack/echo.git ./src/github.com/labstack</code>  Also https://github.com/labstack/gommon.git from Labstack must be imported. Together with: 
https://github.com/labstack/echo.git
https://github.com/labstack/gommon.git

https://github.com/mattn/go-colorable.git
https://github.com/mattn/go-isatty.git
Some of golang extra packages are needed: net, crypto and text. You may get them from: https://github.com/golang add them to the src directory.

You may run the following script for setup: `config.sh`
