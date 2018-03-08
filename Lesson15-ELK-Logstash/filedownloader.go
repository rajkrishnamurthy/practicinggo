package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

//DownloadandSaveFiles : Takes the JSON parsed struct and downloads files from Source URL
func DownloadandSaveFiles(elkdemoinstance Elkdemo) {
	{
		for _, element := range elkdemoinstance.Filesets {

			fmt.Printf("Fileset Name : %v \n", element.Filepersona)
			if element.Action.Download == "yes" { //Download only if true
				success, filename := getFileFromURL(element.Sourceurl, element.Savefileas)
				fmt.Printf("%v, %v", success, filename)
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
