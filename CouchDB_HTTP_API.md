# CouchDB API DOC

_Table of Contents_

[**1. Route** `/`](#1-route)  
[**2. Route** `/_uuids`](#2-route-uuids)  
[**3. Route** `/{db}`](#3-route-db)  
[**4. Route** `/_backup/{db}`](#4-route-backupdb)  
[**5. Route** `/{db}/{docid}`](#5-route-dbdocid)  
[**6. Route** `/{db}/{docid}/{attachment}`](#6-route-dbdocidattachment)  
[**7. Route** `/_replicate`](#7-route-replicate)  
[**8. Route** `/{db}/_compact`](#8-route-dbcompact)  
[**9. Route** `/{db}/_purge`](#9-route-dbpurge)  
[**10. Route** `_all_dbs`](#10-route-alldbs)  

____


## 1.  **Route** `/`

**`GET /`**

**Response**

```http
HTTP/1.1 200 OK
Cache-Control: must-revalidate
Content-Length: 179
Content-Type: application/json
Date: Sat, 10 Aug 2013 06:33:33 GMT
Server: CouchDB (Erlang/OTP)

{
"couchdb": "Welcome",
"uuid": "85fb71bf700c17267fef77535820e371",
"vendor": {
    "name": "The Apache Software Foundation",
    "version": "1.3.1"
},
"version": "1.3.1"
}
```

____

## 2.  **Route** `/_uuids`

**`GET /_uuids`**

**Response**

```http
HTTP/1.1 200 OK
Content-Length: 362
Content-Type: application/json
Date: Sat, 10 Aug 2013 11:46:25 GMT
ETag: "DGRWWQFLUDWN5MRKSLKQ425XV"
Expires: Fri, 01 Jan 1990 00:00:00 GMT
Pragma: no-cache
Server: CouchDB (Erlang/OTP)

{
    "uuids": [
        "75480ca477454894678e22eec6002413",
        "75480ca477454894678e22eec600250b",
        "75480ca477454894678e22eec6002c41",
        "75480ca477454894678e22eec6003b90",
        "75480ca477454894678e22eec6003fca",
        "75480ca477454894678e22eec6004bef",
        "75480ca477454894678e22eec600528f",
        "75480ca477454894678e22eec6005e0b",
        "75480ca477454894678e22eec6006158",
        "75480ca477454894678e22eec6006161"
    ]
}
```

____

## 3. **Route** /{db}

**`HEAD /^[a-z][a-z0-9_$()+/-]*$`**

Database operations - check existence

**Response**

```http
 HTTP/1.1 200 OK
 Server: CouchDB/1.6.1 (Erlang OTP/R16B02)
 Date: Tue, 19 Jan 2016 12:34:30 GMT
 Content-Type: text/plain; charset=utf-8
 Content-Length: 240
 Cache-Control: must-revalidate


 HTTP/1.1 404 Object Not Found
 Server: CouchDB/1.6.1 (Erlang OTP/R16B02)
 Date: Tue, 19 Jan 2016 12:37:44 GMT
 Content-Type: text/plain; charset=utf-8
 Content-Length: 44
 Cache-Control: must-revalidate
```

**`PUT /^[a-z][a-z0-9_$()+/-]*$`**

Database operations - creating a database.

**Response**

```http
 HTTP/1.1 201 Created
 Cache-Control: must-revalidate
 Content-Length: 12
 Content-Type: application/json
 Date: Mon, 12 Aug 2013 08:01:45 GMT
 Location: http://localhost:5984/db
 Server: CouchDB (Erlang/OTP)

 {
    "ok": true
 }

 HTTP/1.1 412 Precondition Failed
 Cache-Control: must-revalidate
 Content-Length: 95
 Content-Type: application/json
 Date: Mon, 12 Aug 2013 08:01:16 GMT
 Server: CouchDB (Erlang/OTP)

 {
    "error": "file_exists",
    "reason": "The database could not be created, the file already exists."
 }

 HTTP/1.1 400 Bad Request
 Cache-Control: must-revalidate
 Content-Length: 194
 Content-Type: application/json
 Date: Mon, 12 Aug 2013 08:02:10 GMT
 Server: CouchDB (Erlang/OTP)

 {
    "error": "illegal_database_name",
    "reason": "Name: '_db'. Only lowercase characters (a-z), digits (0-9), and any of the characters _, $, (, ), +, -, and / are allowed. Must begin with a letter."
 }
```

**`POST /^[a-z][a-z0-9_$()+/-]*$`**

Database operations - create a document with an ID assigned by the database

**Response**

```http
 HTTP/1.1 201 Created
 Cache-Control: must-revalidate
 Content-Length: 95
 Content-Type: application/json
 Date: Tue, 13 Aug 2013 15:19:25 GMT
 ETag: "1-9c65296036141e575d32ba9c034dd3ee"
 Location: http://localhost:5984/db/ab39fe0993049b84cfa81acd6ebad09d
 Server: CouchDB (Erlang/OTP)

 {
    "id": "ab39fe0993049b84cfa81acd6ebad09d",
    "ok": true,
    "rev": "1-9c65296036141e575d32ba9c034dd3ee"
 }

 HTTP/1.1 404 Object Not Found
 Server: CouchDB/1.6.1 (Erlang OTP/R16B02)
 Date: Wed, 20 Jan 2016 09:18:46 GMT
 Content-Type: text/plain; charset=utf-8
 Content-Length: 44
 Cache-Control: must-revalidate

 {
	"error":"not_found",
	"reason":"no_db_file"
 }

 HTTP/1.1 400 Bad Request
 Server: CouchDB/1.6.1 (Erlang OTP/R16B02)
 Date: Wed, 20 Jan 2016 09:22:23 GMT
 Content-Type: text/plain; charset=utf-8
 Content-Length: 196
 Cache-Control: must-revalidate

 {
	"error":"illegal_database_name",
	"reason":"Name: '3pufi'. Only lowercase characters (a-z), digits (0-9), and any of the characters _, $, (, ), +, -, and / are allowed. Must begin with a letter."
 }

 HTTP/1.1 409 Conflict
 Server: CouchDB/1.6.1 (Erlang OTP/R16B02)
 Date: Wed, 20 Jan 2016 09:29:49 GMT
 Content-Type: text/plain; charset=utf-8
 Content-Length: 58
 Cache-Control: must-revalidate

 {
	"error":"conflict",
	"reason":"Document update conflict."
 }
```

**`DELETE /^[a-z][a-z0-9_$()+/-]*$`**

Database operations - delete a database

**Response**

```http
 HTTP/1.1 200 OK
 Cache-Control: must-revalidate
 Content-Length: 12
 Content-Type: application/json
 Date: Mon, 12 Aug 2013 08:54:00 GMT
 Server: CouchDB (Erlang/OTP)

 {
    "ok": true
 }

 HTTP/1.1 400 Bad Request
 Server: CouchDB/1.6.1 (Erlang OTP/R16B02)
 Date: Wed, 20 Jan 2016 14:10:00 GMT
 Content-Type: text/plain; charset=utf-8
 Content-Length: 133
 Cache-Control: must-revalidate

 {
	"error":"bad_request",
	"reason":"You tried to DELETE a database with a ?rev= parameter. Did you mean to DELETE a document instead?"
 }

 HTTP/1.1 404 Object Not Found
 Server: CouchDB/1.6.1 (Erlang OTP/R16B02)
 Date: Wed, 20 Jan 2016 14:12:21 GMT
 Content-Type: text/plain; charset=utf-8
 Content-Length: 41
 Cache-Control: must-revalidate

 {
	"error":"not_found",
	"reason":"missing"
 }
```

**`GET /^[a-z][a-z0-9_$()+/-]*$`**

Database operations - get full information about an existing database

**Response**

```http
 HTTP/1.1 200 OK
 Cache-Control: must-revalidate
 Content-Length: 258
 Content-Type: application/json
 Date: Mon, 12 Aug 2013 01:38:57 GMT
 Server: CouchDB (Erlang/OTP)

 {
    "committed_update_seq": 292786,
    "compact_running": false,
    "data_size": 65031503,
    "db_name": "receipts",
    "disk_format_version": 6,
    "disk_size": 137433211,
    "doc_count": 6146,
    "doc_del_count": 64637,
    "instance_start_time": "1376269325408900",
    "purge_seq": 0,
    "update_seq": 292786
 }
```

____

## 4. **Route** `/_backup/{db}`

**`GET /_backup/^[a-z][a-z0-9_$()+/-]*$`**

Get the database file - a physical backup.

**Response**

```http
 HTTP/1.1 200 OK
 Content-Disposition: attachment; filename="test.bd"
 Content-Length: 24576
 Content-Type: application/octet-stream
 Server: Gouch (Go)
 Date: Wed, 20 Jan 2016 15:40:51 GMT

[message-body; type:application/octet-stream, size:24576 bytes]


 HTTP/1.1 404 Not Found
 Cache-Control: must-revalidate
 Content-Type: application/json; charset=utf-8
 Server: Gouch (Go)
 Date: Wed, 20 Jan 2016 15:39:52 GMT
 Content-Length: 55

 {
	"error":"not_found",
	"reason":"Missing database file!"
 }
```

## 5. **Route** `/{db}/{docid}`

Document operations, no attachments

**`HEAD /{db}/{docid}`**

Document operations - get basic information about the document if it exists.

**Response**

```http
TODO

```

**`GET /{db}/{docid}`**

Document operations - get full document, without attachments

**Response**

```http
TODO

```


**`PUT /{db}/{docid}`**

Document operations - create a new document with the given ID

**Response**

```http
TODO

```

**`DELETE /{db}/{docid}`**

Document operations - delete a document with ID and REV

**Response**

```http
TODO

```

## 6. **Route** `/{db}/{docid}/{attachment}`

**`HEAD  /{db}/{docid}/{attachment}`**


**Response**

```http
TODO

```



**`GET  /{db}/{docid}/{attachment}`**


**Response**

```http
TODO

```


**`PUT  /{db}/{docid}/{attachment}`**

**Response**

```http
TODO

```

**`DELETE  /{db}/{docid}/{attachment}`**


**Response**

```http
TODO

```

## 7. **Route** `/_replicate`

**`POST /_replicate`**

**Response**

```http
TODO

```

## 8. **Route** `/{db}/_compact`

**`POST /{db}/_compact`**

**Response**

```http
TODO

```

## 9. **Route** `/{db}/_purge`

**`POST /{db}/_purge`**

**Response**

```http
TODO

```

## 10. **Route** `_all_dbs`

**`GET /_add_dbs`**

**Response**

```http
  HTTP/1.1 200 OK
  Cache-Control: must-revalidate
  Content-Length: 52
  Content-Type: application/json
  Date: Sat, 10 Aug 2013 06:57:48 GMT
  Server: CouchDB (Erlang/OTP)

  [
   "_users",
   "contacts",
   "docs",
   "invoices",
   "locations"
  ]
```

____