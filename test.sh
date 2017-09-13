#!/bin/bash

#
# Test with curl the database API
#

PROTOCOL=http
ADDRESS=localhost
PORT=5335

URL=$PROTOCOL"://"$ADDRESS":"$PORT
TESTCOUNT=0

# GET /
echo -e "\n\n$TESTCOUNT >>>  GET /\n"
curl -v --noproxy localhost, $URL"/"
let TESTCOUNT++

# GET /_uuids
echo -e "\n\n$TESTCOUNT >>> GET /_uuids\n"
curl -v --noproxy localhost, $URL"/_uuids"
let TESTCOUNT++

# PUT /:{db}
echo -e "\n\n$TESTCOUNT >>> PUT /:{db}\n"
curl -v --noproxy localhost, -X PUT $URL"/test_db"
let TESTCOUNT++

# GET /_all_dbs
echo -e "\n\n$TESTCOUNT >>> /_all_dbs\n"
curl -v --noproxy localhost, $URL"/_all_dbs"
let TESTCOUNT++

# HEAD /:{db}
echo -e "\n\n$TESTCOUNT >>> HEAD /:{db}\n"
curl -v --noproxy localhost, -I HEAD $URL"/test_db"
let TESTCOUNT++

# POST /:{db}
echo -e "\n\n$TESTCOUNT >>> POST /:{db}\n"
curl -v --noproxy localhost, -H "Content-Type: application/json" -d '{"_id":"test_doc" ,"name":"test document"}' $URL"/test_db"
let TESTCOUNT++

# GET /_backup/:{db}
echo -e "\n\n$TESTCOUNT >>> GET /_backup/:{db}\n"
curl -v --noproxy localhost, $URL"/_backup/test_db"
let TESTCOUNT++

# DELETE /:{db}
echo -e "\n\n$TESTCOUNT >>> DELETE /:{db}\n"
curl -v --noproxy localhost, -X DELETE $URL"/test_db"
let TESTCOUNT++


echo -e "\n\n<<< ALL $((TESTCOUNT-1)) TESTS DONE!\n"