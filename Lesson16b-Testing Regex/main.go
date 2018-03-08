package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func main() {

	// //var filename = "logstash.conf"
	// var filename = "newfile"
	// var regexpression string
	// //var sourcestring = "elasticsearch  { \n \t hosts => [\"localhosts:9200\"] \n \t indexvalue => ./nycfile \n }"
	// //var sourcestring = "elasticsearch  { \n hosts=>[\"localhosts:9200\"] \n indexvalue => ./nycfile \n }"
	// fr, err := ioutil.ReadFile(filename)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// //var sourcestring string

	// regexpression = "hosts.*\\[.*\\]"
	// //regexpression = "elasticsearch{1}[\\s\\S]*\\}"
	// //regexpression = "(?m)^.*$"
	// sourcestring := string(fr)

	// re := regexp.MustCompile(regexpression)
	// allstrings := re.FindAllString(sourcestring, -1)
	// fmt.Printf("Array Length = %v \n", len(allstrings))
	// for _, thisstring := range allstrings {
	// 	fmt.Println(thisstring)
	// }

	// fmt.Printf("Reformatted String \n : %s", re.ReplaceAllString(sourcestring, "elasticsearch { }"))
	cleanuptempfiles("./", "\\.archive\\.ao\\.")
}

// filepath should always have a trailing "/". No time to add padding code :)
func cleanuptempfiles(filepath string, filepattern string) bool {

	var filename, matchstring = "", ""
	var err error
	var fileinfoarray []os.FileInfo
	var fileinfo os.FileInfo

	regex := regexp.MustCompile(filepattern)
	fileinfoarray, err = ioutil.ReadDir(filepath)
	if err != nil {
		log.Fatal(err)
		return false
	}

	for _, fileinfo = range fileinfoarray {
		filename = fileinfo.Name()
		matchstring = regex.FindString(filename)
		if matchstring != "" {
			err = os.Remove(filepath + filename)
			if err != nil {
				fmt.Printf("Cannot remove file: %v @ path : %v \n", filename, filepath)
			}
		}
	}
	return true

}
