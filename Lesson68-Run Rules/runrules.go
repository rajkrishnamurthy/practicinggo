//usr/bin/env go run "$0" "$@"; exit "$?"
// All calls come here from bash. This is the main router for all scripts written in Go
package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

var errorNotification = color.New(color.FgRed, color.Bold, color.BlinkSlow).PrintFunc()

var baseFolder = "../"
var commandStrings []string

type ruleType struct {
	ruleHandles       []string
	ruleNamingPattern string
	ruleBaseDir       string
	criteriaPattern   string
	criteria          struct {
		minNum int
		maxNum int
	}
}

type loginType struct {
	loginHost     string
	loginUser     string
	loginPassword string
}

type ruleRunner struct {
	*ruleType
	*loginType
}

var rulerunner *ruleRunner
var rulerangeregex, rulenameregex *regexp.Regexp

func init() {
	// color.NoColor = true
	// cyanunderline := color.New(color.FgCyan).Add(color.Underline).PrintlnFunc()

}

func main() {
	const ruleRange = "200-10200"
	// get the method to be called. this is passed in argument [1]
	// os.Args[1]
	rulerunner := &ruleRunner{}
	rulerunner.loginType = &loginType{
		loginHost:     "cnreverseproxy",
		loginUser:     "raj@fabrikam.com",
		loginPassword: "\\$C0nt1Nub3123",
	}
	rulerunner.ruleType = &ruleType{
		ruleBaseDir:       "/home/rajkrishnamurthy/coding/continube/Flows/rules",
		ruleNamingPattern: `^rule([0-9]+)_`,
		criteriaPattern:   `^([0-9]*)-([0-9]*)$`,
	}
	rulerangeregex = regexp.MustCompile(rulerunner.criteriaPattern)
	rulenameregex = regexp.MustCompile(rulerunner.ruleNamingPattern)

	loginCommands := []string{fmt.Sprintf("continube -i login -H %s %s %s",
		rulerunner.loginHost, rulerunner.loginUser, rulerunner.loginPassword)}
	rulerunner.executeCommands(loginCommands)

	if err := rulerunner.buildRuleCriteria(ruleRange); err != nil {
		log.Fatal(err)
	}
	commandStrings, err := rulerunner.buildCommands()
	if err != nil {
		log.Fatal(err)
	}
	rulerunner.executeCommands(commandStrings)
}

func (rulerunner *ruleRunner) buildCommands() (outputCommands []string, err error) {
	folderElements, err := ioutil.ReadDir(rulerunner.ruleBaseDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, folderElement := range folderElements {
		if folderElement.IsDir() {
			ruleHandle := folderElement.Name()
			fmt.Println("ruleHandle\t", ruleHandle)
			selected, err := rulerunner.checkRuleCriteria(ruleHandle)
			if err != nil {
				continue
			}
			if selected {
				fqrulepath := func(ruleHandle, ruleFolder string) string {
					if strings.TrimSpace(ruleFolder) == "" {
						return ruleHandle
					}
					_, err := os.Stat(ruleFolder + "/" + ruleHandle)
					if err != nil {
						return ""
					}
					return ruleFolder + "/" + ruleHandle
				}(ruleHandle, rulerunner.ruleBaseDir)

				rulerunner.ruleHandles = append(rulerunner.ruleHandles, fqrulepath)
				outputCommands = append(outputCommands,
					fmt.Sprintf("continube -i validate rule %s", fqrulepath))

			}

		}
	}

	outputCommands = append(outputCommands)

	return outputCommands, nil
}

func (rulerunner *ruleRunner) executeCommands(commands []string) (err error) {

	var cmd *exec.Cmd
	for _, command := range commands {
		fmt.Println(command)
		cmd = exec.Command("bash", "-c", command)
		cmd.Env = append(os.Environ())
		var stdoutBuf, stderrBuf bytes.Buffer
		cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
		cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

		err := cmd.Run()
		if err != nil {
			log.Printf("Failed: %s\n", err)
		}

	}
	return nil
}

func (rulerunner *ruleRunner) buildRuleCriteria(ruleRange string) (err error) {
	source := fmt.Sprintf("%s", strings.ToLower(ruleRange))
	matchslice := rulerangeregex.FindAllStringSubmatch(source, -1)
	switch {
	case len(matchslice) == 1:
		// Match found
		submatches := matchslice[0]
		switch {
		case len(submatches) == 3:
			// 3 members; [0] is full match, [1]=min, [2]=max
			if rulerunner.criteria.minNum, err = strconv.Atoi(submatches[1]); err != nil {
				return err
			}
			if rulerunner.criteria.maxNum, err = strconv.Atoi(submatches[2]); err != nil {
				return err
			}
		case len(submatches) == 2:
			// 2 members; [0] is full match, [1]=min or max; determined if "-" is prefix or suffix
		}
	case len(matchslice) == 0:
		return fmt.Errorf("%s", "No Match Found")
	}

	return nil
}

func (rulerunner *ruleRunner) checkRuleCriteria(ruleHandle string) (selected bool, err error) {
	source := fmt.Sprintf("%s", strings.ToLower(ruleHandle))
	matchslice := rulenameregex.FindAllStringSubmatch(source, -1)
	switch {
	case len(matchslice) == 1:
		// Match found
		submatches := matchslice[0]
		switch {
		case len(submatches) == 2:
			// 3 members; [0] is full match, [1]=rule number
			ruleNumber, err := strconv.Atoi(submatches[1])
			if err != nil {
				return false, err
			}
			if ruleNumber >= rulerunner.criteria.minNum && ruleNumber <= rulerunner.criteria.maxNum {
				return true, nil
			}
		case len(submatches) == 1:
			// 2 members; [0] is full match, [1]=min or max; determined if "-" is prefix or suffix
			return false, nil
		}
	case len(matchslice) == 0:
		return false, fmt.Errorf("%s", "No Match Found")
	}
	return false, nil
}
