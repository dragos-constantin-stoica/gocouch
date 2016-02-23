# gocouch
CouchDB implementation in Go. The main target platform is for mobile devices.  
The name is not final, it represents a concatenation of Go and Couch. We are still thinking to call it Gr0_0uchO DB, and add a mustache on a CouchDB logo.  
This project is not another implementation of CouchDB system. It came as an alternative to existing implementations: PouchDB, CouchBase Lite, TouchDB etc for mobile devices.
The main purpose is to allow a decent storage mechanism of data on mobile devices that will be able to replicate/sync with a CouchDB sever. 
By its nature GocouchDB runs off-line, stores data in JSON format and exposes a REST API to the application, and most important it syncs.

# 0. Introduction
Get this repository from github: `git clone https://github.com/iqcouch/gocouch.git`
Setup environment variable: `GOPATH`.  
The storage is based on BoltDB (https://github.com/boltdb/bolt). You need to add the project from github to src directory.  
The REST API is assured by echo framework (http://labstack.com/echo and https://github.com/labstack/echo). Add this framework from github.  Also https://github.com/labstack/gommon.git from Labstack must be imported. Together with: https://github.com/mattn/go-colorable.git and  https://github.com/mattn/go-isatty.git  
Some of golang extra packages are needed: net, crypto and text. You may get them from: https://github.com/golang add them to the src directory.

Or after cloning the repository onto your local drive, `cd gocouch/gocouch; chmod 0777 ./config.sh` you may run the following script for setup: `config.sh`  
If you get this message: `INFO|echo|%s GO CouchDB started ... everybody relax, NOW!` then you are up and running, congratulations!
