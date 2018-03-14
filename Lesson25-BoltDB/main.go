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

const IntType, StringType, TextType, ByteArray, StructType byte = 101, 102, 103, 104, 105

type KV struct {
	Key     string
	KeyType byte
	Value   interface{}
}

func upsertRecords(db *bolt.DB, bktname []byte, kva []KV) (recordsInserted int, err error) {

	var val []byte
	var buf bytes.Buffer
	var commitTx int

	for _, kv := range kva {
		if kv.Value == nil {
			continue
		}
		key := []byte(kv.Key)
		enc := gob.NewEncoder(&buf)
		if err = enc.Encode(kv); err != nil {
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

	return commitTx, nil
}

func selectRecords(db *bolt.DB, bktname []byte, key string) (kva []KV, err error) {
	var buf bytes.Buffer
	var kvtmp KV

	err = db.View(func(tx *bolt.Tx) (err error) {
		var k []byte
		bkt := tx.Bucket([]byte(bktname))

		// If the user passes a specific key, the kv Array will be 1 x 1
		if key != "" {
			dec := gob.NewDecoder(&buf)
			k = []byte(key)
			n, err := buf.Read(bkt.Get(k))
			if err != nil {
				fmt.Printf("number of bytes read %v \n", n)
				return err
			}
			err = dec.Decode(&kvtmp)
			if err != nil {
				return err
			}
			kva = append(kva, kvtmp)
			buf.Reset()
		} else { // If the users does not pass a key value (i.e, ""), we will create a n x 1 array
			cur := bkt.Cursor()
			for k, kv := cur.First(); k != nil; k, kv = cur.Next() {
				rdr := bytes.NewReader(kv)
				n, err := buf.ReadFrom(rdr)
				dec := gob.NewDecoder(&buf)
				if err != nil {
					fmt.Printf("number of bytes read %v \n", n)
					return err
				}
				err = dec.Decode(&kvtmp)
				if err != nil {
					return err
				}
				kva = append(kva, kvtmp)
				buf.Reset()
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

	kva := make([]KV, 1, 1)
	kva = append(kva, KV{"key4", IntType, 54321})
	kva = append(kva, KV{"key5", StringType, "Testing Key5"})
	kva = append(kva, KV{"key6", TextType, `Updating: hello
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
	}
	fmt.Printf("Selected Records %v \n", kvaa)

}
