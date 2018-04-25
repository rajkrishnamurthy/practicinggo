package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

const (
	hostname = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "password"
	dbname   = "tempNode"
)

func main() {
	connectionstring := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, password, hostname, port, dbname, "disable")
	db, err := gorm.Open("postgres", connectionstring)
	if err != nil {
		log.Printf("gorm.Open() Error = %v \n", err)
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	if !db.HasTable(&WorkerNode{}) {
		db.AutoMigrate(&WorkerNode{})
		fmt.Printf("AutoMigrate() complete \n")

		// Create a record
		db.Create(&WorkerNode{
			CNAPIVersion:  "1.0.0",
			CNPath:        "/usr/lib",
			CNExternalURL: "http://node1.continube.io",
		})
		fmt.Printf("Create() complete \n")

	} else {
		fmt.Printf("Model already exists! \n")
	}

	// // Read
	// var product Product
	// db.First(&product, 1)                   // find product with id 1
	// db.First(&product, "code = ?", "L1212") // find product with code l1212

	// // Update - update product's price to 2000
	// db.Model(&product).Update("Price", 2000)
	// fmt.Printf("Model().Update() complete \n")

	// // Delete - delete product
	// //db.Delete(&product)
}
