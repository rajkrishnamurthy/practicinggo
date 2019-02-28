package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	var err error

	inst := new(TaskInstance)
	inputs := &Inputs{}
	outputs := &Outputs{}

	// inputs.WorkingPath = "F:\\temp"
	inputs.InputFileName = "https://bit.ly/2UeLftY"
	inputs.OCRHeaderKey = "2cba1b29da294f7395b2819220bd3e03"
	inputs.OCRQString = "language=en"
	inputs.OCRURL = "https://southeastasia.api.cognitive.microsoft.com/vision/v1.0/ocr"

	if err = inst.OCRProcessor(inputs, outputs); err != nil {
		fmt.Printf("Error = %v \n", err)
	}
	fmt.Printf("Print Output Struct \n %s", outputs.OCROutput)
}

// OCRProcessor : Outputs OCR lines based on a text
// This function is registered through a TCP handler
// RPC Client Call
func (inst *TaskInstance) OCRProcessor(inputs *Inputs, outputs *Outputs) (err error) {
	var client *http.Client
	var urlString string
	var outputBytesArray []byte

	type OCRParser struct {
		Language    string  `json:"language"`
		Orientation string  `json:"orientation"`
		TextAngle   float64 `json:"textAngle"`
		Regions     []struct {
			BoundingBox string `json:"boundingBox"`
			Lines       []struct {
				BoundingBox string `json:"boundingBox"`
				Words       []struct {
					BoundingBox string `json:"boundingBox"`
					Text        string `json:"text"`
				} `json:"words"`
			} `json:"lines"`
		} `json:"regions"`
	}

	ocrparser := OCRParser{}

	// if _, err = os.Stat(inputs.WorkingPath); err != nil {
	// 	outputs.OCROutput = []byte(fmt.Sprintf("Working Path not valid. %v", err))
	// 	errDesc := "Working Path not valid. %v"
	// 	outputs.OCROutput = []byte(fmt.Sprintf(errDesc, err))
	// 	return fmt.Errorf(errDesc, err)

	// }
	// if err = os.Chdir(inputs.WorkingPath); err != nil {
	// 	errDesc := "Cannot change directory. %v"
	// 	outputs.OCROutput = []byte(fmt.Sprintf(errDesc, err))
	// 	return fmt.Errorf(errDesc, err)
	// }

	// if _, err = os.Stat(inputs.InputFileName); err != nil {
	// 	errDesc := "File not found. %v"
	// 	outputs.OCROutput = []byte(fmt.Sprintf(errDesc, err))
	// 	return fmt.Errorf(errDesc, err)
	// }

	err = DownloadFile("./imagefile", inputs.InputFileName)
	if err != nil {
		errDesc := "Cannot download file. %v"
		outputs.OCROutput = []byte(fmt.Sprintf(errDesc, err))
		return fmt.Errorf(errDesc, err)
	}

	fileByteArray, err := ioutil.ReadFile("./imagefile")
	if err != nil {
		errDesc := "Cannot read file. %v"
		outputs.OCROutput = []byte(fmt.Sprintf(errDesc, err))
		return fmt.Errorf(errDesc, err)
	}
	bytesReader := bytes.NewReader(fileByteArray)

	urlString = inputs.OCRURL + "?" + inputs.OCRQString
	client = &http.Client{}

	req, err := http.NewRequest("POST", urlString, bytesReader)
	req.Header.Add("Ocp-Apim-Subscription-Key", inputs.OCRHeaderKey)
	req.Header.Add("Content-Type", "application/octet-stream")
	resp, err := client.Do(req)
	if err != nil {
		errDesc := "Cannot send request. %v"
		outputs.OCROutput = []byte(fmt.Sprintf(errDesc, err))
		return fmt.Errorf(errDesc, err)
	}
	defer resp.Body.Close()

	outputBytesArray, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		errDesc := "No body in response. %v"
		outputs.OCROutput = []byte(fmt.Sprintf(errDesc, err))
		return fmt.Errorf(errDesc, err)
	}
	fmt.Printf("%s", outputBytesArray)

	if err = json.Unmarshal(outputBytesArray, &ocrparser); err != nil {
		errDesc := "Cannot unmarshal response. %v"
		outputs.OCROutput = []byte(fmt.Sprintf(errDesc, err))
		return fmt.Errorf(errDesc, err)

	}

	var wordColl string
	for _, region := range ocrparser.Regions {
		for _, line := range region.Lines {
			for _, word := range line.Words {
				wordColl = wordColl + strings.TrimSpace(word.Text)
				// fmt.Printf("text = %s", word.Text)
			}
			wordColl = wordColl + "\n"
		}
	}
	// outputs.OCROutput = []byte(fmt.Sprintf("%s", ocrparser))
	outputs.OCROutput = []byte(wordColl)

	return nil
}

type TaskInstance int

type Inputs struct {
	WorkingPath   string
	InputFileName string
	OCRURL        string // The URL to access OCR processing
	OCRQString    string // Query String to be appended to URL
	OCRHeaderKey  string // Header Key for OCR
}

type Outputs struct {
	OCROutput []byte // OCR Output encoded as []byte
}

// DownloadFile : Download the file
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
