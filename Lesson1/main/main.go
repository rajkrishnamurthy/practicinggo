package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/boltdb/bolt"
)

// Objective : Can we pass a pointer to an array as a parameter to a function

// var Testvar1, testvar2 int = 1, 20

func passtoFunction(a []int) int {
	// b := *a
	var b []int
	b = make([]int, 5)
	copy(b, a)
	summer := 0
	for value := range b {
		summer += value
	}
	b[0] = 100
	return summer
}

func main() {
	//var i int = 10
	// for i := Testvar1; i < testvar2; i++ {
	// 	fmt.Printf("Test Variable Value = %v \n", i)
	// }
	var parma []int
	var parmb *[]int

	parma = []int{1, 2, 3, 4, 5}
	parmb = &parma
	fmt.Printf("Paramemter A: %v \n Parameter B: %v \n Return Value = %v \n", parma, parmb, passtoFunction(parma))

	{
		fileinfo, err := os.Stat("main.go")
		fmt.Printf("os.Stat Fileinfo = %v \n Error = %v \n", fileinfo, err)
	}
	{
		dirPath, err := os.Getwd()
		fmt.Printf("Directory = %v \n Error = %v \n", dirPath, err)
	}
	callBoltDB()
}

func callBoltDB() {
	var boltFile = "snippets.boltdb"
	// var path = "F:\\Coding\\GoProg\\OutsidetheCourse\\PracticingGo\\Lesson1\\main"
	var path = ".\\"
	db, err := bolt.Open(filepath.Join(path, boltFile), 0644, nil)
	if err != nil {
		fmt.Printf("Error : %v \n", err)
	}
	defer db.Close()
	err = db.View(func(tx *bolt.Tx) error {
		return tx.ForEach(func(name []byte, _ *bolt.Bucket) error {
			fmt.Println(string(name))
			return nil
		})
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	err = db.View(func(tx *bolt.Tx) error {
		var bkterr error
		bkt := tx.Bucket([]byte("SnippetsByID"))
		bkterr = bkt.ForEach(func(k, v []byte) error {
			fmt.Printf("A %v is %v.\n", k, v)
			concatstring := string(k) + "\n" + string(v) + "\n"
			if err := ioutil.WriteFile("outputfile.dmp", []byte(concatstring), os.ModeAppend); err != nil {
				fmt.Printf("Error with file operations")
			}
			return nil
		})
		if bkterr != nil {
			fmt.Printf("Error inside fetching bucket values \n")
			return bkterr
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
}
