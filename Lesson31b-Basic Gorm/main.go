package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	hostname = "localhost"
	port     = "5432"
	username = "postgres"
	password = "password"
	dbname   = "temp"
)

// User has and belongs to many languages, use `user_languages` as join table
type User struct {
	// gorm.Model
	UserID       int `sql:",pk" gorm:"AUTO_INCREMENT"`
	UserName     string
	UserLocation string
	LanguageID   uint `gorm:"foreignkey:LanguageID,assoication_foreignkey:LanguageID"`
	// Languages    []Language  `gorm:"foreignkey:ID`// `gorm:"many2many:user_languages;"`
}

type Language struct {
	// gorm.Model
	LanguageID int `sql:",pk" gorm:"AUTO_INCREMENT,PRIMARY_KEY"`
	Name       string
	State      string
	Country    string
}

func main() {

	// usr := User{}
	// language := Language{}
	// languages := make([]Language, 0)

	connectionstring := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", username, password, hostname, port, dbname, "disable")
	db, err := gorm.Open("postgres", connectionstring)
	if err != nil {
		log.Printf("gorm.Open() Error = %v \n", err)
		panic("failed to connect database")
	}
	defer db.Close()

	db.DropTableIfExists(&Language{})
	db.CreateTable(&Language{})
	db.DropTableIfExists(&User{})
	db.CreateTable(&User{})

	db.Model(&User{}).Related(&Language{})
	//// SELECT * FROM "languages" INNER JOIN "user_languages" ON "user_languages"."language_id" = "languages"."id" WHERE "user_languages"."user_id" = 111

}
