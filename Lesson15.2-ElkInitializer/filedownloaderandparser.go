package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

//DownloadandSaveFiles : Takes the JSON parsed struct and downloads files from Source URL
func DownloadSaveandSubsetFiles(elkdemoinstance Elkdemo) {
	{
		for _, element := range elkdemoinstance.Filesets {

			fmt.Printf("Fileset Name : %v \n", element.Filepersona)
			if element.Action.Download == "yes" { //Download only if true
				success, filename := getFileFromURL(element.Sourceurl, element.Savefileas)
				fmt.Printf("%v, %v", success, filename)
			}
		}
	}

	{
		for _, element := range elkdemoinstance.Filesets {

			fmt.Printf("Fileset Name : %v \n", element.Filepersona)
			if element.Action.Subset == "yes" { //Subsetfiles
				success, bytesw := SubsetCSVFile(element.Savefileas, element.Action.Subsetnumber)
				fmt.Printf("Successful? : %v, and Bytes Written %v", success, bytesw)
			}
		}
	}

}

func getFileFromURL(url string, outputfile string) (bool, string) {
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]
	fmt.Println("Downloading", url, "to", fileName)

	if strings.TrimSpace(outputfile) != "" {
		fileName = outputfile
	}

	// TODO: check file existence first with io.IsExist
	output, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error while creating", fileName, "-", err)
		return false, ""
	}
	defer output.Close()

	//tr := http.DefaultTransport.(*http.Transport)
	//tr.TLSClientConfig.InsecureSkipVerify = true

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		log.Fatal(err)
		return false, ""
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return false, ""
	}

	fmt.Println(n, "bytes downloaded.")
	return true, fileName
}

//SubsetCSVFile : Passes the Input Filename and Subset String as Inputs and return Success flag and Bytes Written as Outputs
func SubsetCSVFile(filename string, subsetnumber string) (bool, int) {

	var recordset, pcts []string
	var filesize = 0
	var pcti, counter, pctp = 0, 0, 0.0
	var err error
	var subsetpct, firsttimerecord = false, true

	if strings.Contains(subsetnumber, "%") {
		subsetpct = true
		pcts = strings.Split(subsetnumber, "%")
		if pcts[0] != "" {
			pcti, err = strconv.Atoi(pcts[0])
			if err != nil {
				fmt.Printf("Cannot convert pct to int %v \n", err)
				log.Fatal(err)
				return false, 0
			}
			pctp = float64(pcti) / 100.00
			fmt.Printf("%v", pctp)
		} else {
			pcti, err = strconv.Atoi(subsetnumber)
			if err != nil {
				fmt.Printf("Cannot convert to int %v \n", err)
				log.Fatal(err)
				return false, 0
			}

		}
	}

	fileinfo, err := os.Stat(filename)
	if err != nil {
		fmt.Println("Error Opening CSV File", err)
		log.Fatal(err)
		return false, 0
	}

	filesize = int(fileinfo.Size())
	fmt.Printf("%d", filesize)

	csvfile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error Opening CSV File", err)
		log.Fatal(err)
		return false, 0
	}

	reader := bufio.NewReader(csvfile)
	for {
		record, err := reader.ReadString('\n')
		if firsttimerecord {
			if subsetpct {
				pcti = int(pctp * float64(filesize) / float64(len(record)))
				firsttimerecord = false
			}
		}
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println("Error Reading CSV File", err)
				log.Fatal(err)
				return false, 0
			}
		}
		recordset = append(recordset, record)
		counter++

		if counter > pcti { //indicated value is actual number of records
			break
		}

	}
	csvfile.Close()

	success, bytesw := writetofile("new", recordset, filename)
	if !success {
		return false, 0
	}
	return true, bytesw
}
