package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

type kv struct {
	key   string
	value interface{}
}

func createBucket(dbname string) (db *bolt.DB, bkt *bolt.Bucket, tx *bolt.Tx) {
	//var err error
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	tx, err = db.Begin(true)
	if err != nil {
		log.Fatal(err)
		return nil, nil, nil
	}

	bkt, err = tx.CreateBucket([]byte("test"))
	if err != nil {
		log.Fatal(err)
		return nil, nil, nil
	}

	return db, bkt, tx
}

func createTx(bkt *bolt.Bucket, tx *bolt.Tx, kva []kv) (commitTx int) {

	for _, kv := range kva {

		key := []byte(kv.key)
		if val, ok := kv.value.([]byte); ok {
			if err := bkt.Put(key, val); err != nil {
				log.Fatal(err)
				tx.Rollback()
			}

		}

		if err := tx.Commit(); err != nil {
			log.Fatal("Rollback failed")
		}
		commitTx++
	}
	return commitTx

}

func main() {
	db, bkt, tx := createBucket("my.db")
	defer db.Close()

	kva := make([]kv, 10, 10)
	kva = append(kva, kv{"key1", 12345})
	kva = append(kva, kv{"key2", "value2"})
	kva = append(kva, kv{"key1", `hello 
		is this alright string
		let us test
		`})

	successfulTx := createTx(bkt, tx, kva)
	fmt.Printf("Number of Committed Transactions %v \n", successfulTx)

}
