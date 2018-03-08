package main

import (
	"encoding/json"
	"io/ioutil"
)

type Fileaction struct {
	Download     string `json:"download"`
	Update       string `json:"update"`
	Extract      string `json:"extract"`
	Subset       string `json:"subset"`
	Subsetnumber string `json:"subsetnumber"`
	Otherinfo    string `json:"otherinfo"`
}

type Fileset struct {
	Filepersona string	`json:"filepersona"`
	Sourceurl   string     `json:"sourceurl"`
	Description string     `json:"description"`
	Filetype    string     `json:"filetype"`
	Savefileas  string     `json:"savefileas"`
	Action      Fileaction `json:"action"`
}

type Patternaction string

type Patterner struct {
	Name         string        `json:"name"`
	Order        int           `json:"order"`
	Description  string        `json:"description"`
	Action       Patternaction `json:"action"`
	Outputfile   string        `json:"outputfile"`
	Findpattern  string        `json:"findpattern"`
	Replacevalue string        `json:"replacevalue"`
	Filenames    []string      `json:"filenames"`
}

type Elkdemo struct {
	Name            string      `json:"name"`
	Filesets        []Fileset   `json:"filesets"`
	Searchandupdate []Patterner `json:"searchandupdate"`
}

func DecodeJSONFile(pfilename string) (bool, Elkdemo) {

	var elkdemo01 Elkdemo
	filecontents, err := ioutil.ReadFile(pfilename)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(filecontents, &elkdemo01)
	//fmt.Printf("%v \n", elkdemo01)
	return true, elkdemo01
}
