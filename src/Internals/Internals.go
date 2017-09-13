// Internals project Internals.go
package Internals

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	_ "fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
	"github.com/labstack/echo"
)

const (
	WelcomeMsg = "GO CouchDB started ... everybody relax, NOW!"
	ServerMsg  = "GouchDB (Go)"
)

func GetMD5Hash(text []byte) string {
	hash := md5.Sum(text)
	return hex.EncodeToString(hash[:])
}

func GetUUID() string {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		//w.Write(err)
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return hex.EncodeToString(uuid)
}

type BoltDB struct {
	db *bolt.DB
}

func NewBoltDB(filepath string, dbname string) *BoltDB {
	db, err := bolt.Open(filepath+dbname, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}

	return &BoltDB{db}
}

func ROBoltDB(filepath string, dbname string) *BoltDB {
	db, err := bolt.Open(filepath+dbname, 0666, &bolt.Options{ReadOnly: true})
	if err != nil {
		log.Fatal(err)
	}

	return &BoltDB{db}
}

func (b *BoltDB) Path() string {
	return b.db.Path()
}

func (b *BoltDB) Close() {
	b.db.Close()
}

func (b *BoltDB) Read(bucket []byte, key []byte) (value []byte, err error) {

	err = b.db.View(func(tx *bolt.Tx) error {
		bucket_tmp := tx.Bucket(bucket)
		value = bucket_tmp.Get(key)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (b *BoltDB) UpdateDB(bucket []byte, key []byte, value []byte) error {
	return b.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			log.Fatal("Error creating bucket: %s", err)
			return err
		}
		err = b.Put(key, value)
		if err != nil {
			log.Fatal("Error writing to bucket: %s", err)
		}
		return err
	})
}

func (b *BoltDB) ExistsDoc(bucket []byte) bool {
	var result = false
	b.db.View(func(tx *bolt.Tx) error {
		bucket_tmp := tx.Bucket(bucket)
		result = bucket_tmp != nil
		return nil
	})
	return result
}

func DeleteDB(filepath string, dbname string) error {
	return os.Remove(filepath + dbname)
}

func (b *BoltDB) ExportFile(dbname string, c echo.Context) error {
	err := b.db.View(func(tx *bolt.Tx) error {
		c.Response().Header().Set(echo.HeaderContentType, "application/octet-stream")
		c.Response().Header().Set("Content-Disposition", `attachment; filename="`+dbname+`"`)
		c.Response().Header().Set("Content-Length", strconv.Itoa(int(tx.Size())))
		_, err := tx.WriteTo(c.Response().Writer())
		return err
	})
	return err
}
