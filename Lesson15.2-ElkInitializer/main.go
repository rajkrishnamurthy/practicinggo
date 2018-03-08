package main

import (
	"fmt"
)

/*
Purpose: Dynamically push the dataset into Logstash. This program runs on the Logstash container instance
Author: Raj Krishnamurthy
Parameters, in default order:
	a. -dataset : Source Dataset Location. An HTTP URL or a fully qualified file path for dataset
	b. -conf: Source Logstash Configuration File. An HTTP URL or a fully qualified file path for Logstash.conf
	c. -eserver: Elastic Server Info. Name or IP address of the elastic server and port name [format: <server ip/name>:<port>]
	d. (optional) -inputyaml: An HTTP URL or a fully qualified file path for YAML

Output:

Outline of Steps:
Step1: Setup a "Flag" Infrastructure
Step2: Parse Flags for Source URL/Filepath
Step3: Valdate Flags & Exception Handling
Step4: Fetch Source Files: Dataset and Config
Step5: Parse and Update Configuration file with ElasticSearch Server
Step6: Set Environment Variables
Step7: Run Logstash

References:
Using Flags: https://gobyexample.com/command-line-flags
File Download: https://github.com/thbar/golang-playground/blob/master/download-files.go
Fetch https: https://stackoverflow.com/questions/12122159/golang-how-to-do-a-https-request-with-bad-certificate

Commandline: Example: ./Lesson15-ELK-Logstash.exe -dataset="http://dataset" -conf="http://conf" 1 2 3
./Lesson15-ELK-Logstash.exe -dataset="https://data.cityofnewyork.us/api/views/h9gi-nx95/rows.csv" -conf="https://raw.githubusercontent.com/rajkrishnamurthy/examples/master/Exploring%20Public%20Datasets/nyc_traffic_accidents/nyc_collision_logstash.conf"


*/
func main() {

	_, elkdemo01 := DecodeJSONFile("elk_demo2.0.json")
	{
		for _, element := range elkdemo01.Filesets {

			fmt.Printf("Fileset Name : %v \n", element.Filepersona)
		}
	}
	DownloadSaveandSubsetFiles(elkdemo01)

	fmt.Printf("Download and Save and Subset Files Section Complete \n")
	{
		for _, element := range elkdemo01.Searchandupdate {
			fmt.Printf("Element Name: %v \n", element.Name)
		}
	}

	statusarray := SearchandUpdateFiles(elkdemo01)
	for _, statuslineitem := range statusarray {
		fmt.Printf("Section : %v \t, Error Description: %v \n", statuslineitem.section, statuslineitem.errorstring)
	}
}
