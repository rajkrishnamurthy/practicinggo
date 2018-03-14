package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

var dbname string = "test.db"
var bktname []byte = []byte("bucket1")

type kv struct {
	key   string
	value interface{}
}

func upsertRecords(db *bolt.DB, bktname []byte, kva []kv) (recordsInserted int, err error) {

	var val []byte
	var buf bytes.Buffer
	var commitTx int

	for _, kv := range kva {
		if kv.value == nil {
			continue
		}
		key := []byte(kv.key)
		enc := gob.NewEncoder(&buf)
		if err = enc.Encode(kv.value); err != nil {
			log.Fatal(err)
			return 0, err
		}
		val = buf.Bytes()

		err = db.Update(func(tx *bolt.Tx) (err error) {
			bkt, err := tx.CreateBucketIfNotExists(bktname)
			if err != nil {
				log.Fatal(err)
				return err
			}

			if err = bkt.Put(key, val); err != nil {
				log.Fatal(err)
				return err
			}
			fmt.Printf("Inside createTx : Key = %v \t Value = %v \n", key, val)

			commitTx++
			buf.Reset()
			return nil
		})

		if err != nil {
			log.Fatal(err)
			return 0, err
		}
	}

	return 0, nil
}

func selectRecords(db *bolt.DB, bktname []byte, key string) (kva []kv, err error) {
	err = db.View(func(tx *bolt.Tx) (err error) {
		var k, v []byte
		bkt := tx.Bucket([]byte(bktname))

		if key != "" {
			k = []byte(key)
			v = bkt.Get(k)
			kva = append(kva, kv{string(k), v})
		} else {
			cur := bkt.Cursor()
			for k, v = cur.First(); k != nil; k, v = cur.Next() {
				kva = append(kva, kv{string(k), v})
			}
		}
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return kva, nil
}

func main() {

	db, err := bolt.Open(dbname, 0444, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	kva := make([]kv, 1, 1)
	kva = append(kva, kv{"key4", 54321})
	kva = append(kva, kv{"key5", "Testing Key5"})
	kva = append(kva, kv{"key6", `Updating: hello
			is this alright string
			let us test
			`})
	fmt.Printf("%v \n", kva)

	cnt, err := upsertRecords(db, bktname, kva)
	fmt.Printf("Inserted Records %v \n", cnt)
	if err != nil {
		log.Fatal(err)
		return
	}

	kvaa, err := selectRecords(db, bktname, "")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Selected Records %v \n", kvaa)

}
