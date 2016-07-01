:bangbang: This code represents **Work In Progress!** Do not use it in production environment. :bangbang:

# gocouch
CouchDB implementation in Go. The main target platform is for mobile devices.  
The name is not final, it represents a concatenation of Go and Couch. We are still thinking to call it Gr0_0uchO DB, and add a mustache on a CouchDB logo. Another names on the list: Recamier, SquaBD, seTTes (a palindromized form of settee and pronounced [si tɪts]), Banquette , all refering to couch in English. The tanslation/equivalent word in other languages: gogol, bedi DB, MoengaDB (just to annoy mongoDB fans), mag-abang DB.  
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
If you get this message: `INFO|echo| GO CouchDB started ... everybody relax, NOW!` then you are up and running, congratulations!

# 1. Development Status
The following REST API are implemented or on the roadmap:  
- [x] `GET http://server:5984/ `
- [x] `GET http://server:5984/_uuids `
- [x] `GET http://server:5984/_all_dbs `
- [x] `HEAD http://server:5984/{db} `
- [x] `GET http://server:5984/{db} ` -- Partially implemented. The statistics is not recorder. It will be fully implemented as part of replication mechanism.
- [x] `POST http://server:5984/{db} `
- [x] `PUT http://server:5984/{db} `
- [x] `DELETE http://server:5984/{db} `
- [x] `GET http://server:5984/_backup/{db} ` -- This is specific API and allows you to download the associate BoltDB file. Like a full dump of the database.
- [ ] `POST http://server:5984/_replicate `
- [ ] `GET http://server:5984/{db}/_local/{doc} `
- [ ] `POST http://server:5984/{db}/_local/{doc} `
- [ ] `GET http://server:5984/{db}/_changes `
- [ ] `POST http://server:5984/{db}/_changes ` 
- [ ] `POST http://server:5984/{db}/_revs_diff ` 
- [ ] `GET http://server:5984/{db}/{doc} ` 
- [ ] `PUT http://server:5984/{db}/{doc} ` 
- [ ] `POST http://server:5984/{db}/_bulk_docs ` 
- [ ] `POST http://server:5984/{db}/_ensure_full_commit `  

Database log is directed to the console in this moment but this will be changed for production release, it will most probably be directed toward a file or a database. There is a config database `_gcfg.bd`, in a future version it will be an API that will allow configuration management. Attachments are also on the roadmap, both as document attachments and as CouchApps.


Information sources
====

* http://docs.couchdb.org/en/stable/replication/protocol.html
* https://git-wip-us.apache.org/repos/asf?p=couchdb.git;a=blob;f=src/couch_replicator/src/couch_replicator_utils.erl;h=d7778db;hb=HEAD
* https://media.readthedocs.org/pdf/couchdb/latest/couchdb.pdf
* http://dataprotocols.org/couchdb-replication/
* https://github.com/rcouch/rcouch/wiki/Replication-Algorithm
* http://www.replication.io/databases
 

