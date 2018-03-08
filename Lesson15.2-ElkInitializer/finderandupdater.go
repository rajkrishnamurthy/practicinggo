package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type LineItemStatus struct {
	status      bool
	section     string // indicates json processor section
	errorstring string // details error
	objectinfo  interface{}
}

var archivecounter = 0

//SearchandUpdateFiles : Takes the JSON parsed struct, regex Find Pattern and Inserts/Replaces Values
func SearchandUpdateFiles(elkdemoinstance Elkdemo) []LineItemStatus {
	var statusarray []LineItemStatus
	{
		var elementactionstring, foundstring = "", ""
		for _, element := range elkdemoinstance.Searchandupdate {
			fmt.Printf("Element Name : %v \n", element.Name)
			elementactionstring = string(element.Action)
			// Is there a "find" clause in action string

			if strings.Contains(elementactionstring, "find") {
				var filecontents, inputfilename, findpattern, outputfilename, replacestring, printstring string = "", "", "", "", "", ""
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
				foundstring = regex.FindString(filecontents)
				printstring = foundstring
				//Merge will have a filename in its replace value indicated by f:///<filename>
				var tempreplacestring []string
				var mergesourcefilename string
				if strings.Contains(elementactionstring, "merge") {
					tempreplacestring = strings.Split(replacestring, "f:///")
					if tempreplacestring[0] == "" { // I thought that [0] will have the match characteristics. No. It is blank
						mergesourcefilename = tempreplacestring[1]
						{
							fileinfo, err := os.Stat(mergesourcefilename)
							fmt.Printf("File Name and Size : %v, %v \n", fileinfo.Name(), fileinfo.Size())
							if err != nil {
								log.Fatal(err)
								{
									status := LineItemStatus{false, "Searchandupdate", "Merge file does not exist", element}
									statusarray = append(statusarray, status)
								}
								fmt.Printf("Merge file does not exist \n")
								return statusarray
							}
						}

						{
							bytes, err := ioutil.ReadFile(mergesourcefilename)
							if err != nil {
								log.Fatal(err)
								{
									status := LineItemStatus{false, "Searchandupdate", "Merge file cannot be read", element}
									statusarray = append(statusarray, status)
								}
								fmt.Printf("Merge file cannot be read \n")
								return statusarray
							}
							replacestring = string(bytes)
						}
					}

				}

				// insert and replace are only a subset of find
				if strings.Contains(elementactionstring, "replace") || strings.Contains(elementactionstring, "merge") {
					filecontents = regex.ReplaceAllString(filecontents, replacestring)
					printstring = filecontents
				}

				// TODO: Need to get this section done
				// how do i know where to insert generically
				if strings.Contains(elementactionstring, "insert") {
					if len(foundstring) == 0 {
						//section incomplete
					} else {
						{
							status := LineItemStatus{false, "Searchandupdate", "Pattern match found. But Insert flag also found", element}
							statusarray = append(statusarray, status)
						}

					}
				}
				fmt.Printf("Input Filename \n %v \n Printstring %v", inputfilename, printstring)
				{
					var namesakearray []string
					namesakearray = append(namesakearray, printstring) // TODO : array created unnecessarily. just kept intact for compatibility with other calls. need to revise later
					writetofile("new", namesakearray, outputfilename)
				}
			} //if action = 'find' ends

		}
	}
	cleanuptempfiles("./", "\\.archive\\.ao\\.")
	return statusarray
}

func writetofile(fileoperation string, writestrings []string, outputfilename string) (bool, int) {
	//fileop can be one of the following:
	//	a. new -> create a new file
	//	b. append -> add to existing file, create new one if it does not exist
	var err error
	var bytesw int
	var fh *os.File
	var tempoutputfilename = "tempfilename" + strconv.Itoa(rand.Intn(20))
	var filealreadyexists, finallyswapfiles = false, false

	_, err = os.Stat(outputfilename)
	if err == nil {
		filealreadyexists = true
	}

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

	for _, writestring := range writestrings {
		bytesw, err = fh.WriteString(writestring)
		if err != nil {
			log.Fatal(err)
			return false, 0
		}
	}
	fh.Sync()
	fh.Close()
	fmt.Printf("Variable [ignore] %v \n", finallyswapfiles)

	if finallyswapfiles {
		swapFiles(tempoutputfilename, outputfilename)
	}

	return true, bytesw
}

func swapFiles(sourcefile string, targetfile string) bool {
	var err error

	_, err = os.Stat(sourcefile)
	if err != nil {
		log.Fatal(err)
		fmt.Printf("%v File does not exist. Error Description %v \n", sourcefile, err)
		return false
	}

	_, err = os.Stat(sourcefile)
	if err != nil {
		log.Fatal(err)
		fmt.Printf("%v File does not exist. Error Description %v \n", sourcefile, err)
		return false
	}

	err = os.Rename(targetfile, targetfile+".archive."+"ao."+strconv.Itoa(archivecounter))
	if err != nil {
		log.Fatal(err)
		fmt.Printf("Eror in Renaming File (with archive) %v", err)
		return false
	}
	archivecounter++

	// err = os.Remove(outputfilename)
	// if err != nil {
	// 	log.Fatal(err)
	// 	fmt.Printf("Eror in Removing File %v", err)
	// 	return false, 0
	// }
	os.Rename(sourcefile, targetfile)
	if err != nil {
		log.Fatal(err)
		fmt.Printf("Eror in Renaming File (to original name) %v", err)
		return false
	}

	return true

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
