package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
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
type paramstype map[string]interface{}

func setFlags() [][]string {
	//setFlags creates a 2D, n x 3 array with
	//	a. [0] = Flag Name
	//	b. [1] = Flag Default Value. File Name for Downloads, Section Value for Section, Replace String for Updates
	//	c. [2] = Flag Description
	//	d. [3] = Action: download | update | section
	//	e. [4] = Save File As for Downloads or Find String for Updates

	var mx2paramlist [][]string
	var mx1paramlist []string

	//TODO: To be uncommented. Large file. Avoided during testing
	mx1paramlist = []string{"dataset",
		"https://data.cityofnewyork.us/api/views/h9gi-nx95/rows.csv",
		"Data injested into Elastic Search", "download", "dataset.csv"}
	mx2paramlist = append(mx2paramlist, mx1paramlist)

	mx1paramlist = []string{"logstashconf",
		"https://raw.githubusercontent.com/rajkrishnamurthy/examples/master/Exploring%20Public%20Datasets/nyc_traffic_accidents/nyc_collision_logstash.conf",
		"Field Mapping Configuration for Logstash",
		"download", "logstash.conf"}
	mx2paramlist = append(mx2paramlist, mx1paramlist)

	mx1paramlist = []string{"elasticconf",
		"https://raw.githubusercontent.com/elastic/examples/master/Exploring%20Public%20Datasets/nyc_traffic_accidents/nyc_collision_template.json",
		"Rules Mapping Configuration for ElasticSearch",
		"download", "template.json"}
	mx2paramlist = append(mx2paramlist, mx1paramlist)

	mx1paramlist = []string{"kibanaconf",
		"https://raw.githubusercontent.com/elastic/examples/master/Exploring%20Public%20Datasets/nyc_traffic_accidents/nyc_collision_kibana.json",
		"Configuration File for Kibana",
		"download", "kibana.json"}
	mx2paramlist = append(mx2paramlist, mx1paramlist)

	mx1paramlist = []string{"section", "elasticsearch{1}[\\s\\S]*\\}", "Section that will be used for UPDATES", "update", "none"}
	mx2paramlist = append(mx2paramlist, mx1paramlist)

	//Provided space to align
	mx1paramlist = []string{"elasticserver", "hosts => [\"elasticserver01:9200\"] \n",
		"Allows Dynamic Configuration of Elastic Search Server",
		"update", "hosts.*\\[.*\\]"}
	mx2paramlist = append(mx2paramlist, mx1paramlist)

	mx1paramlist = []string{"index", "index => \"nyc_visionzero\" \n",
		"Specifies the Elastic Search Index Name",
		"update", "index.*"}
	mx2paramlist = append(mx2paramlist, mx1paramlist)

	mx1paramlist = []string{"template", "template => \"./template.json\" \n",
		"Elasticsearch Template JSON file that needs to be downloaded",
		"update", "template[^_].*"}
	mx2paramlist = append(mx2paramlist, mx1paramlist)

	mx1paramlist = []string{"templatename", "template_name => \"nyc_visionzero\" \n",
		"Elasticsearch Template Name",
		"update", "template_name.*"}
	mx2paramlist = append(mx2paramlist, mx1paramlist)

	mx1paramlist = []string{"args", "none", "Arguments passed outside flags in Commandline", "none", "none"}
	mx2paramlist = append(mx2paramlist, mx1paramlist)

	return mx2paramlist
}

// Step2
//This function will take the "flags" array as input parameter and return a map w/ parsed output
func assignFlags(pFlags [][]string) paramstype {
	//pFlags is a 2D, n x 5 array with
	//	a. [0] = Flag Name
	//	b. [1] = Flag Default Value or Regex Pattern
	//	c. [2] = Flag Description
	//	d. [3] = Action Flag
	//	e. [4] = Save File As

	tempparmnames := make(paramstype)

	for _, pkey := range pFlags {
		// Initialize all Inteface{} values to nil
		tempparmnames[pkey[0]] = flag.String(pkey[0], pkey[1], pkey[2])
	}

	flag.Parse()

	//case "args": Args() is assigned post Parse()
	tempparmnames["args"] = flag.Args()

	return tempparmnames

}

//Step4-a
//This function takes a Slice of all files that need to be fetched, temp disk location and returns error, if any
func fetchFiles(paramflags paramstype, paramlist [][]string) bool {
	//TODO : Error Handling
	//TODO : Output file/path override
	var tempfetchlist [][]string

	{
		var downloadfile, savefileas string
		var flagvalueindex = 0
		var downloadflagindex = 3
		var savefileasindex = 4

		var temp1xd []string

		{
			for _, pvalue := range paramlist {
				downloadfile = ""
				if strings.Contains(strings.ToLower(pvalue[downloadflagindex]), "download") {
					downloadfile = strings.TrimSpace(*paramflags[pvalue[flagvalueindex]].(*string))
					savefileas = strings.TrimSpace(pvalue[savefileasindex])
					if downloadfile != "" {
						temp1xd = []string{downloadfile, savefileas}
						//add file to tempfetchlist
						tempfetchlist = append(tempfetchlist, temp1xd)
					}
				}

			}
		}
		{
			for _, pvalue := range tempfetchlist {
				downloadFromURL(pvalue[0], pvalue[1])
			}
		}
	}
	return true
}

//Step4-b
func downloadFromURL(url string, outputfile string) bool {
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
		return false
	}
	defer output.Close()

	//tr := http.DefaultTransport.(*http.Transport)
	//tr.TLSClientConfig.InsecureSkipVerify = true

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		log.Fatal(err)
		return false
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return false
	}

	fmt.Println(n, "bytes downloaded.")
	return true
}

//This is the parent function that replaces the old parse, delete and insert functions
//This is based on RegEx :). I accept!!
func processFiles(basefilename string, paramflags paramstype,
	paramlist [][]string) (rresultval bool, rmodfilename string, rbyteswritten int) {
	//TODO : Error Handling
	//TODO : Output file/path override
	var tempfetchlist [][]string
	var modifiedfilestring string
	//The key assumption is that there is only one "Section" per execution
	var section string = strings.TrimSpace(*paramflags["section"].(*string))
	var modfilename = "newfile"

	{
		var replacestring, findstring string
		//[0] contains the commandline parameter to replace the attribute
		var attributereplaceindex = 0

		//[3] indicates whether the record is "update" or "download" or "section"
		var updateflagindex = 3

		//[4] contains the pre-set search pattern for the flag
		var findstringindex = 4

		var temp1xd []string

		{
			for _, pvalue := range paramlist {
				replacestring = ""
				//Only select those records that have the "update" flag set
				if strings.Contains(strings.ToLower(pvalue[updateflagindex]), "update") {
					replacestring = strings.TrimSpace(*paramflags[pvalue[attributereplaceindex]].(*string))
					findstring = strings.TrimSpace(pvalue[findstringindex])
					if replacestring != "" && pvalue[attributereplaceindex] != "section" {
						temp1xd = []string{findstring, replacestring}
						//add file to tempfetchlist
						tempfetchlist = append(tempfetchlist, temp1xd)
					}
				}

			}
		}
		_, modifiedfilestring = configFinalizer(basefilename, section, tempfetchlist)
		fmt.Printf("New File at Parent Function \n %s", modifiedfilestring)
	}

	fw, err := os.Create(modfilename)
	if err != nil {
		log.Fatal(err)
	}
	defer fw.Close()
	fwwriter := bufio.NewWriter(fw)
	bytesw, werr := fwwriter.WriteString(modifiedfilestring)
	if werr != nil {
		log.Fatal(werr)
	}
	fwwriter.Flush()
	fw.Sync()
	return true, modfilename, bytesw
}

func configFinalizer(filename string, sectionregex string,
	findreplacearray [][]string) (bool, string) {
	//var outputbool = false
	var filecontents, newfilecontents, sectioncontents, newsectioncontents = "", "", "", ""
	var regexouter, regexinner *regexp.Regexp

	fr, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	filecontents = string(fr)

	//Should we use the section string as regex or just the name
	//for example; "elasticsearch{1}[\\s\\S]*\\}"
	regexouter = regexp.MustCompile(sectionregex)
	allstrings := regexouter.FindAllString(filecontents, -1)
	if len(allstrings) > 1 {
		fmt.Printf("Multiple Matches. Array Length = %v \n", len(allstrings))
		return false, string(len(allstrings)) + " matches found"
	} else if len(allstrings) == 1 {
		sectioncontents = allstrings[0]
	} else if len(allstrings) < 1 {
		return false, "no matches found"
	}

	newsectioncontents = sectioncontents
	var findregex, replacestring = "", "'"
	//we have the section contents. iterate through the array to find/replace attributes
	//findregex should be a regex. Example; "hosts.*\\[.*\\]\""
	//replacestring should be of the form = "hosts => \"[elasticsearch01:9200]\" \n"

	for _, pvalue := range findreplacearray {
		findregex, replacestring = pvalue[0], pvalue[1]

		regexinner = regexp.MustCompile(findregex)
		allstrings := regexinner.FindAllString(newsectioncontents, -1)
		for _, thisstring := range allstrings {
			fmt.Printf("Search found: %s \n", thisstring)
			newsectioncontents = regexinner.ReplaceAllString(newsectioncontents, replacestring)
			fmt.Printf("Replaced: %s \n", newsectioncontents)
		}

	}

	fmt.Printf("Modifed Section Contents in Child Function \n %s", newsectioncontents)

	//
	newfilecontents = regexouter.ReplaceAllString(filecontents, newsectioncontents)

	fmt.Printf("Modifed File Contents in Child Function \n %s", newfilecontents)

	return true, newfilecontents
}

func main() {
	/*
		//Commandline Args
		// -dataset=<http URL for dataset> : File will be named as dataset.csv
		// -logstashconf=<http URL for > : File will be named as logstash.conf
		// -elasticconf=<http URL for .JSON file > : File will be named as template.json
		// -kibanaconf=<http URL for Kibana JSON file> : File will be named as kibana.json
		// -section=elasticsearch : Has a preset expression "elasticsearch{1}[\\s\\S]*\\}"
		// -elasticserver="hosts => [\"elasticserver01:9200\"] \n" : Regex expression that updates elasticsearch configuration in logstash.conf

		//Step1: Setup a Flag Infrastructure
		paramlist := setFlags()
		//paramlist is a 2d n x 3 matrix
		//paramlist := []string{"dataset", "logstashconf", "elasticconf", "eserver", "inputyaml", "args"}

		paramnames := make(paramstype)
		paramnames = assignFlags(paramlist)

		for _, value := range paramlist {
			switch value[0] {
			case "args":
				//No action
			default:
				fmt.Printf("%s : %s \n", value[0], *paramnames[value[0]].(*string))
			}

		}

		if fetchFiles(paramnames, paramlist) {
			fmt.Println("successful downloads")
		}

		_, newfilename, byteswritten := processFiles("logstash.conf", paramnames, paramlist)
		fmt.Printf("Finally: Filename %s \t Byteswritten : %d \n", newfilename, byteswritten)

		//testingfunction()
	*/
	_, elkdemo01 := DecodeJSONFile("elk_demo2.0.json")
	{
		for _, element := range elkdemo01.Filesets {

			fmt.Printf("Fileset Name : %v \n", element.Filepersona)
		}
	}
	{
		for _, element := range elkdemo01.Searchandupdate {
			fmt.Printf("Element Name: %v \n", element.Name)
		}
	}
}
