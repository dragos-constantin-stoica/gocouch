# gocouch
CouchDB implementation in Go. The main target platform is for mobile devices.  
The name is not final, it represents a concatenation of Go and Couch. We are still thinking to call it Gr0_0uchO DB, and add a mustache on a CouchDB logo.  
This project is not another implementation of CouchDB system. It came as an alternative to existing implementations: PouchDB, CouchBase Lite, TouchDB etc for mobile devices.
The main purpose is to allow a decent storage mechanism of data on mobile devices that will be able to replicate/sync with a CouchDB sever. 
By its nature GocouchDB runs off-line, stores data in JSON format and exposes a REST API to the application, and most important it syncs.

# 0. Introduction
The storage is based on BoltDB (). You need to add the project from github  
The REST API is assured by echo framework (). Add this framework from github  
Some of golang extra packages are needed: ... . Add them to the src directory by   
