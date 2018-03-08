package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

type LineItemStatus struct {
	status      bool
	section     string // indicates json processor section
	errorstring string // details error
	objectinfo  interface{}
}

//SearchandUpdateFiles : Takes the JSON parsed struct, regex Find Pattern and Inserts/Replaces Values
func SearchandUpdateFiles(elkdemoinstance Elkdemo) []LineItemStatus {
	var statusarray []LineItemStatus
	{
		var elementactionstring string
		var allstrings []string

		for _, element := range elkdemoinstance.Searchandupdate {
			fmt.Printf("Element Name : %v \n", element.Name)
			elementactionstring = string(element.Action)
			// Is there a "find" clause in action string
			if strings.Contains(elementactionstring, "find") {
				var inputfilename, findpattern, outputfilename, replacestring string = "", "", "", ""
				var filecontents = ""
				var regex *regexp.Regexp

				inputfilename = element.Filenames[0] //right now the program does not accept multiple input filenames eventhough defined in the jsonfile as an array
				findpattern = element.Findpattern
				outputfilename = element.Outputfile
				replacestring = element.Replacevalue

				fr, err := ioutil.ReadFile(inputfilename)
				if err != nil {
					log.Fatal(err)
					{
						status := LineItemStatus{false, "Searchandupdate", "cannot read input file in find", element}
						statusarray = append(statusarray, status)
					}
					return statusarray
				}
				filecontents = string(fr)

				regex = regexp.MustCompile(findpattern)
				allstrings = regex.FindAllString(filecontents, -1)
				writetofile("new", allstrings, outputfilename)
				// insert and replace are only a subset of find
				if strings.Contains(elementactionstring, "replace") {
					filecontents = regex.ReplaceAllString(filecontents, replacestring)
					allstrings = append(allstrings, filecontents)
					writetofile("new", allstrings, outputfilename)
				}
				if strings.Contains(elementactionstring, "insert") {
					if len(allstrings) == 0 {
						// how do i know where to insert generically

					} else {
						{
							status := LineItemStatus{false, "Searchandupdate", "Pattern match found. But Insert flag also found", element}
							statusarray = append(statusarray, status)
						}

					}
				}

			} //if action = 'find' ends

		}
	}
	return statusarray
}

func writetofile(fileoperation string, writestrings []string, outputfilename string) (bool, int) {
	//fileop can be one of the following:
	//	a. new -> create a new file
	//	b. append -> add to existing file, create new one if it does not exist
	var err error
	var bytesw int
	var fh *os.File
	var fileinfo os.FileInfo
	var tempoutputfilename = "tempfilename"
	var filealreadyexists, finallyswapfiles = false, false

	fileinfo, err = os.Stat(outputfilename)
	if fileinfo.Name() != "" && fileinfo.Size() > 0 {
		filealreadyexists = true
	}
	// if err != nil {
	// 	log.Fatal(err)
	// 	return false, 0
	// }

	if fileoperation == "new" {
		if !filealreadyexists {
			fh, err = os.Create(outputfilename)
			if err != nil {
				log.Fatal(err)
				return false, 0
			}
		} else {
			fh, err = os.Create(tempoutputfilename)
			if err != nil {
				log.Fatal(err)
				return false, 0
			}
			finallyswapfiles = true
		}
	} else if fileoperation == "append" {
		if !filealreadyexists {
			fh, err = os.Create(outputfilename)
			if err != nil {
				log.Fatal(err)
				return false, 0
			}
		} else {
			fh, err = os.OpenFile(outputfilename, os.O_APPEND|os.O_RDWR, 777)
			if err != nil {
				log.Fatal(err)
				return false, 0
			}
		}
	}
	defer fh.Close()

	for _, writestring := range writestrings {
		bytesw, err = fh.WriteString(writestring)
		if err != nil {
			log.Fatal(err)
			return false, 0
		}
	}
	fh.Sync()

	if finallyswapfiles {
		os.Rename(outputfilename, outputfilename+".archive")
		os.Remove(outputfilename)
		os.Rename(tempoutputfilename, outputfilename)
	}
	return true, bytesw
}
