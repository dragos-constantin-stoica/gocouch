// all_dbs.go
package main

import (
	_"log"
	"path/filepath"
)

var (
	ALL_DBS []string
)



func Contains(list []string, elem string) bool { 
        for _, t := range list { if t == elem { return true } } 
        return false 
} 

//Look for *.bd files in the current location
//Store this information in ALL_DBS variable
//Use this variable as a proxy/cache
func all_dbs_init(bdpath string) {
	ALL_DBS = []string{}
	files, _ := filepath.Glob(bdpath + "*.bd")
	//remove _replicator and _gcfg from the list
	for _, values := range files {
		//remove .bd extension
		dbName := values[:len(values)-3]
		if ValidDB_Name.MatchString(dbName) || dbName == "_users"{
			ALL_DBS = append(ALL_DBS, dbName)		
		} 
	}
	//ALL_DBS = files
	//log.Println(ALL_DBS)
}
