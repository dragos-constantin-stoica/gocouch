// GConfig project GConfig.go
package GConfig

import (
	"Internals"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type GoCouch struct {
	UUID string `json:"uuid"`
}

type HTTPd struct {
	Bind_Address string `json:"bind_address,omitempty"`
	Port         string `json:"port,omitempty"`
}

//Config structure
//Contains sections
//Each section contains key value pairs
type CFG struct {
	GoCouch `json:"gocouch"`
	HTTPd   `json:"httpd"`
}

const (
	CFG_BD_NAME        = "_gcfg.bd"
	REPLICATOR_BD_NAME = "_replicator.bd"
)

var (
	GoCouchCFG CFG
	cfgDB      *Internals.BoltDB
)

//init - put default config values here
func init() {
	GoCouchCFG.GoCouch.UUID = Internals.GetUUID()
	GoCouchCFG.HTTPd.Bind_Address = "127.0.0.1"
	GoCouchCFG.HTTPd.Port = "5984"

}

// GOCouch configuration is stored in a BoltDB
// the location of the file is _gcfg.bd
func CheckCFGFile() {

	if _, err := os.Stat(CFG_BD_NAME); err != nil {
		fmt.Println(err)
		if os.IsNotExist(err) {
			//create a cfg database and save it for next time
			cfgDB = Internals.NewBoltDB("", CFG_BD_NAME)
			defer cfgDB.Close()
			value, _ := json.Marshal(GoCouchCFG)
			err = cfgDB.UpdateDB([]byte("config"), []byte("key"), value)
			if err != nil {
				log.Fatal(err)
			}

		}
	} else {
		//we have a cfg database; load the data
		cfgDB = Internals.ROBoltDB("", CFG_BD_NAME)
		value, err := cfgDB.Read([]byte("config"), []byte("key"))
		err = json.Unmarshal(value, &GoCouchCFG)
		if err != nil {
			log.Fatal(err)
		}
	}
}

//Check for replicator db file and create one if none
func CheckReplicatorFile() {
	if _, err := os.Stat(REPLICATOR_BD_NAME); err != nil {
		if os.IsNotExist(err) {
			syncDB := Internals.NewBoltDB("", REPLICATOR_BD_NAME)
			defer syncDB.Close()
		}
	}
}
