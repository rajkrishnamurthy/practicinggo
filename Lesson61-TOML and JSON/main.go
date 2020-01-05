package main

import (
	"fmt"
	"log"
	"regexp"

	"github.com/pelletier/go-toml"
)

type appGroupTOML struct {
	configTOMLTree    *toml.Tree
	developerTOMLTree *toml.Tree
	defaultMap        map[string]interface{}
	targetTOMLTree    *toml.Tree
}

const configTOMLFile = "continubeConfigMap.toml"
const developerTOMLFile = "developerInputs.toml"
const defaultTOMLFile = "defaultInputs.toml"
const targetTOMLFile = "targetAppGroup.toml"

const signature = "%developerInputs%."

func main() {

	var err error
	regExpAttr := regexp.MustCompile(signature)

	appGroup := &appGroupTOML{}
	appGroup, err = appGroup.initAppGroupTOML(configTOMLFile, targetTOMLFile)
	if err != nil {
		log.Fatalf("%v", err)
	}

	developerTOMLTree, defaultMap, err := appGroup.initOtherTOMLFiles(developerTOMLFile, defaultTOMLFile)
	if err != nil {
		log.Fatalf("%v", err)
	}

	appGroup.developerTOMLTree = developerTOMLTree
	appGroup.defaultMap = defaultMap

	err = appGroup.constructConfigTree(appGroup.configTOMLTree, regExpAttr)
	if err != nil {
		log.Fatalf("%v", err)
	}

}

func (appGroup *appGroupTOML) initAppGroupTOML(configTOMLFile, targetTOMLFile string) (retAppGroup *appGroupTOML, err error) {
	var configTOMLTree, targetTOMLTree *toml.Tree
	retAppGroup = &appGroupTOML{}
	configTOMLTree, err = toml.LoadFile(configTOMLFile)
	if err != nil {
		return nil, err
	}
	retAppGroup.configTOMLTree = configTOMLTree
	retAppGroup.developerTOMLTree = nil
	retAppGroup.defaultMap = make(map[string]interface{}, 0)

	targetTOMLTree, err = toml.LoadFile(targetTOMLFile)
	if err != nil {
		return nil, err
	}
	retAppGroup.targetTOMLTree = targetTOMLTree

	return retAppGroup, nil
}

func (appGroup *appGroupTOML) initOtherTOMLFiles(developerTOMLFile, defaultTOMLFile string) (developerTOMLTree *toml.Tree, defaultMap map[string]interface{}, err error) {
	var defaultTOMLTree *toml.Tree
	developerTOMLTree, err = toml.LoadFile(developerTOMLFile)
	if err != nil {
		return nil, nil, fmt.Errorf("Cannot fetch developer TOML. %v", err)
	}
	defaultTOMLTree, err = toml.LoadFile(defaultTOMLFile)
	if err != nil {
		return nil, nil, fmt.Errorf("Cannot fetch default TOML. %v", err)
	}
	defaultTOMLMap := defaultTOMLTree.ToMap()

	return developerTOMLTree, defaultTOMLMap, nil

}

func (appGroup *appGroupTOML) constructConfigTree(tomlTree *toml.Tree, regExpAttr *regexp.Regexp) error {
	// Get the top level keys for the tomlTree
	var topLevelKeys []string
	var isString, isNested bool = false, false
	var strVal string
	var intfVal interface{}

	topLevelKeys = tomlTree.Keys()
	if len(topLevelKeys) <= 0 {
		// TOML file is blank. Something is wrong.
		return fmt.Errorf("Configuration File in toml is blank")
	}

	// Check for value substitution
	for _, topLevelKey := range topLevelKeys {
		// Is the topLevelKey a string
		intfVal = tomlTree.Get(topLevelKey)
		appGroup.targetTOMLTree.Set(topLevelKey, intfVal)

		_, isNested = intfVal.(*toml.Tree)
		if isNested {
			continue
		}
		strVal, isString = tomlTree.Get(topLevelKey).(string)
		if !isString {
			// It is not a string
		}
		attrArray := regExpAttr.Split(strVal, -1)
		if len(attrArray) > 1 {
			// This means that there is %developerInputs% match. i.e, this value needs to be substituted
			// Fetch the value from the developerInputs file
			inputPath := attrArray[1] // Always in the 2nd subscript. Can simplify this
			inputArray, err := appGroup.getDeveloperInputs(inputPath)
			if err != nil {
				return fmt.Errorf("%v", err)
			}

			if len(inputArray) == 0 {
				return fmt.Errorf("%s", "No developer inputs present!")
			}

			for _, inputVal := range inputArray {
				if inputVal != "" {
					appGroup.targetTOMLTree.Set(inputPath, inputVal)
				}

			}

		}

	}

	return nil
}

func (appGroup *appGroupTOML) getDeveloperInputs(inputPath string) (inputArray []string, err error) {

	inputArray = []string{}
	defaultInputVal := appGroup.defaultMap[inputPath]

	if devInputString, isString := appGroup.developerTOMLTree.Get(inputPath).(string); isString {
		if devInputString != "" {
			// The developer has provided inputs. So, we should use that
			inputArray = append(inputArray, devInputString)
		} else {
			// The developer has not provided inputs. So, let us try the default map
			if defaultInputString, isDefaultString := defaultInputVal.(string); isDefaultString {
				inputArray = append(inputArray, defaultInputString)
			} else {
				// Report exception
				return inputArray, fmt.Errorf("Values exist neither in input nor in default")
			}
		}

	} else if devInputStringArr, isStringArr := appGroup.developerTOMLTree.Get(inputPath).([]string); isStringArr {
		if len(devInputStringArr) > 0 {
			// The developer has provided inputs. So, we should use that
			inputArray = append(inputArray, devInputStringArr...)
		} else {
			// The developer has not provided inputs. So, let us try the default map
			if defaultInputStringArr, isDefaultStringArr := defaultInputVal.([]string); isDefaultStringArr {
				inputArray = append(inputArray, defaultInputStringArr...)
			} else {
				// Report exception
				return inputArray, fmt.Errorf("Values exist neither in input nor in default")
			}
		}

	} else {
		return inputArray, fmt.Errorf("Unable to fetch developer input string")
	}

	return inputArray, nil
}
