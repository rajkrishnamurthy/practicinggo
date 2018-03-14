package main

import (
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

func main() {

	db, err := bolt.Open(dbname, 0444, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	/*

		kva := make([]kv, 1, 1)
		kva = append(kva, kv{"key4", 12345})
		kva = append(kva, kv{"key5", "value2"})
		kva = append(kva, kv{"key6", `hello
			is this alright string
			let us test
			`})
		fmt.Printf("%v \n", kva)

		_ = db.Update(func(tx *bolt.Tx) (err error) {
			var val []byte
			var buf bytes.Buffer
			var commitTx int

			bkt, err := tx.CreateBucketIfNotExists(bktname)
			if err != nil {
				log.Fatal(err)
				return err
			}

			for _, kv := range kva {
				if kv.value == nil {
					continue
				}
				key := []byte(kv.key)
				enc := gob.NewEncoder(&buf)
				if err = enc.Encode(kv.value); err != nil {
					log.Fatal(err)
					return err
				}
				val = buf.Bytes()

				if err = bkt.Put(key, val); err != nil {
					log.Fatal(err)
					return err
				}
				fmt.Printf("Inside createTx : Key = %v \t Value = %v \n", key, val)

				commitTx++
				buf.Reset()
			}
			return nil
		})

	*/

	_ = db.View(func(tx *bolt.Tx) (err error) {
		fmt.Printf("Inside db.View \n")
		var k, v []byte
		bkt := tx.Bucket([]byte(bktname))
		cur := bkt.Cursor()
		// k = []byte("key1")
		// v = bkt.Get(k)
		// k, v = cur.First()
		for k, v = cur.First(); k != nil; k, v = cur.Next() {
			fmt.Printf("Key = %v \n Value = %v \n \n", string(k[:]), v)
		}
		return nil
	})

}
