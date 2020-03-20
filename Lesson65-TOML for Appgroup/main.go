package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/pelletier/go-toml"
)

type appGroupTOML struct {
	configTOMLTree    *toml.Tree
	developerTOMLTree *toml.Tree
	defaultTOMLTree   *toml.Tree
	targetTOMLTree    *toml.Tree
}

const configTOMLFile = "continubeConfigMap.toml"
const developerTOMLFile = "developerInputs.toml"
const defaultTOMLFile = "defaultInputs.toml"
const targetTOMLFile = "targetAppGroup.toml"

const signature = "%developerInputs%."

func main() {
	// callTestToml("testapp.toml")
	testAppBlueprint()
}

func callTestToml(tomlFile string) {

	configTOMLTree, err := toml.LoadFile(tomlFile)
	if err != nil {
		log.Fatal(err)
	}

	targetTOMLTree, err := toml.LoadFile("targetAppGroup.toml")
	if err != nil {
		log.Fatal(err)
	}

	// {
	// 	iterKey := []string{"apps", "apps.servers"}
	// 	iterVal := configTOMLTree.GetPath(iterKey)
	// 	iterType := strings.ToLower(reflect.TypeOf(iterVal).String())
	// 	fmt.Printf("Get Output: Key=%v\tType=%v\tValue=%v\n", iterKey, iterType, iterVal)
	// }
	// return

	for _, iterKey := range configTOMLTree.Keys() {
		// keypathSlice := iterKey
		iterVal := configTOMLTree.Get(iterKey)
		iterType := strings.ToLower(reflect.TypeOf(iterVal).String())
		fmt.Printf("Get Output: Key=%v\tType=%v\n", iterKey, iterType)
		switch iterType {
		case "string":
			targetTOMLTree.Set(iterKey, iterVal)
			// fmt.Printf("Get Output: Key=%v\tType=%v\tValue=%v\n", iterKey, iterType, iterVal)
			// fmt.Printf("Get Output: Key=%v\tType=%v\n", iterKey, iterType)

		case strings.ToLower("*toml.Tree"): // This is an object
		case strings.ToLower("[]*toml.Tree"): // This is an array of objects
			iterTreeArray, ok := iterVal.([]*toml.Tree)
			if ok {
				for _, iterTree2 := range iterTreeArray {
					for _, iterKey2 := range iterTree2.Keys() {
						// keypathSlice = keypathSlice + "." + iterKey2
						iterVal2 := iterTree2.Get(iterKey2)
						iterType2 := strings.ToLower(reflect.TypeOf(iterVal2).String())
						// fmt.Printf("Get Output: Key=%v\tType=%v\tValue=%v\n", iterKey2, iterType2, iterVal2)
						fmt.Printf("Get Output: Key=%v\tType=%v\n", iterKey2, iterType2)
						targetTOMLTree.Set(iterKey, iterVal)
						// fmt.Printf("Key Path=%s\n", keypathSlice)
					}

				}
			}

		default:

		}

		// fmt.Printf("GetPath Output: Key=%v\tValue=%v\n", iterKey, configTOMLTree.GetPath([]string{iterKey}))
		// fmt.Printf("GetPosition Output: Key=%v\tValue=%v\n", iterKey, configTOMLTree.GetPosition(iterKey))
		// fmt.Printf("GetPositionPath Output: Key=%v\tValue=%v\n", iterKey, configTOMLTree.GetPositionPath([]string{iterKey}))

		strval, _ := targetTOMLTree.ToTomlString()
		fmt.Printf("\n%s\n\n", strval)

	}

	// jsonString, err := json.Marshal(configTOMLTree.ToMap())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(jsonString))

}

func testAppBlueprint() {
	var err error

	appGroup := &appGroupTOML{}
	appGroup, err = appGroup.initTOMLFiles(configTOMLFile, targetTOMLFile, developerTOMLFile, defaultTOMLFile)
	if err != nil {
		log.Fatalf("%v", err)
	}

	appGroup.printAppBlueprint(appGroup.configTOMLTree, "")

	targetTOMLString, err := appGroup.targetTOMLTree.ToTomlString()
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Printf("%s\n", targetTOMLString)

}

func (appGroup *appGroupTOML) printAppBlueprint(tomlTree *toml.Tree, baseKey string) (offsetKey string, err error) {
	topLevelKeys := tomlTree.Keys()
	if len(topLevelKeys) <= 0 {
		return "", fmt.Errorf("%s", "No keys. Is the file blank")
	}
	for _, topLevelKey := range topLevelKeys {
		if baseKey == "" {
			offsetKey = strings.Join([]string{topLevelKey}, "")
		} else {
			offsetKey = strings.Join([]string{baseKey, topLevelKey}, ".")
		}
		// Is the topLevelKey a string
		intfValues := tomlTree.Get(topLevelKey)
		switch strings.ToLower(reflect.TypeOf(intfValues).String()) {

		// This is the leaf element in the tree represented as STRING
		// TODO: This can also be in other data types. So need to check.
		case strings.ToLower("string"):
			valCastToString, _ := intfValues.(string)
			if err := appGroup.writeItemToTargetTOML(valCastToString, offsetKey); err != nil {
				return "", err
			}
			// fmt.Printf("Key=%s\tValue=%s\n", topLevelKey, intfValues)

		// This is the leaf element in the tree represented as []STRING
		// The String array is represented as []interface {}
		// TODO: This can also be in other data []types. So need to check.
		case strings.ToLower("[]interface {}"):
			arrItems, ok := intfValues.([]interface{})
			if ok {
				for arrKey, arrItem := range arrItems {
					valCastToString, _ := arrItem.(string)
					// if err := appGroup.writeItemToTargetTOML(valCastToString, strings.Join([]string{offsetKey, strconv.Itoa(arrKey)}, ".")); err != nil {
					// 	return "", err
					// }
					if err := appGroup.writeItemToTargetTOML(valCastToString, offsetKey+"["+strconv.Itoa(arrKey)+"]"); err != nil {
						return "", err
					}

				}
			}
		case strings.ToLower("*toml.Tree"):
			appGroup.printAppBlueprint(intfValues.(*toml.Tree), offsetKey)
		case strings.ToLower("[]*toml.Tree"):
			arrItems, ok := intfValues.([]*toml.Tree)
			if ok {
				for tomlTreeKey, tomlTreeItem := range arrItems {
					// appGroup.printAppBlueprint(tomlTreeItem, strings.Join([]string{offsetKey, strconv.Itoa(tomlTreeKey)}, "."))
					appGroup.printAppBlueprint(tomlTreeItem, offsetKey+"["+strconv.Itoa(tomlTreeKey)+"]")
				}
			}
		default:
		}

	}
	return offsetKey, nil
}

func (appGroup *appGroupTOML) writeItemToTargetTOML(indexString, PathKey string) error {
	isInput, inputSlice, err := appGroup.checkForDeveloperInputs(fmt.Sprintf("%s", indexString))
	if err != nil {
		return fmt.Errorf("Error in getting Developer Input values")
	}
	if isInput {
		switch {
		case len(inputSlice) == 1: // Not an array of Input values. Only one.
			appGroup.targetTOMLTree.Set(PathKey, inputSlice[0])
		case len(inputSlice) > 1: // Need to create clones of the same level
			for _, iVal := range inputSlice {
				appGroup.targetTOMLTree.Set(PathKey, iVal)
			}

		default:
		}
	} else { // This is not an input match. Just write it to target TOML
		appGroup.targetTOMLTree.Set(PathKey, indexString)
	}

	return nil
}

func (appGroup *appGroupTOML) checkForDeveloperInputs(indexString string) (bool, []string, error) {

	// The format of the indexString is expected as follows
	// %developerInputs%.<keyname>
	// The key can be a string or an []string
	splitStrSlice := strings.SplitAfter(indexString, signature)

	if len(splitStrSlice)-1 < 1 {
		return false, []string{}, nil
	} else if len(splitStrSlice)-1 > 1 {
		return false, []string{}, fmt.Errorf("%s", "Check the Configuration Map. Does not map to valid signature")
	}

	intfValues := appGroup.developerTOMLTree.Get(splitStrSlice[1]) // slile[0] will contain the match %developerInputs%.

	switch strings.ToLower(reflect.TypeOf(intfValues).String()) {
	// This is the leaf element in the tree represented as STRING
	// TODO: This can also be in other data types. So need to check.
	case strings.ToLower("string"):
		return true, []string{fmt.Sprintf("%s", intfValues)}, nil

	// This is the leaf element in the tree represented as []STRING
	// The String array is represented as []interface {}
	// TODO: This can also be in other data []types. So need to check.
	case strings.ToLower("[]interface {}"):
		arrItems, ok := intfValues.([]interface{})
		if ok {
			var retArray = []string{}
			for _, arrItem := range arrItems {
				retArray = append(retArray, fmt.Sprintf("%s", arrItem))
			}
			return true, retArray, nil
		}
	}

	return false, []string{}, nil
}

func callAppBlueprint() {
	var err error
	regExpAttr := regexp.MustCompile(signature)

	appGroup := &appGroupTOML{}
	appGroup, err = appGroup.initTOMLFiles(configTOMLFile, targetTOMLFile, developerTOMLFile, defaultTOMLFile)
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = appGroup.constructConfigTree(appGroup.configTOMLTree, regExpAttr, "")

	tree, isTree := appGroup.targetTOMLTree.Get("apps").([]*toml.Tree)
	if isTree {
		for _, t := range tree {
			fmt.Println(t.Get("appName"))
		}
	}

	fmt.Println(appGroup.targetTOMLTree)
	fmt.Println(appGroup.configTOMLTree)
	m := appGroup.targetTOMLTree.ToMap()
	fmt.Println(appGroup.targetTOMLTree)
	jsonString, _ := json.Marshal(m)
	fmt.Println(string(jsonString))
	if err != nil {
		log.Fatalf("%v", err)
	}

}

func loadTOML(tomlFile string) (*toml.Tree, error) {
	tomlTree, err := toml.LoadFile(tomlFile)
	if err != nil {
		return nil, err
	}
	return tomlTree, nil
}

func (appGroup *appGroupTOML) initTOMLFiles(configTOMLFile, targetTOMLFile, developerTOMLFile, defaultTOMLFile string) (retAppGroup *appGroupTOML, err error) {
	retAppGroup = &appGroupTOML{}

	if retAppGroup.configTOMLTree, err = loadTOML(configTOMLFile); err != nil {
		return nil, err
	}

	if retAppGroup.developerTOMLTree, err = loadTOML(developerTOMLFile); err != nil {
		return nil, err
	}

	if retAppGroup.defaultTOMLTree, err = loadTOML(defaultTOMLFile); err != nil {
		return nil, err
	}

	if retAppGroup.targetTOMLTree, err = loadTOML(targetTOMLFile); err != nil {
		return nil, err
	}

	return retAppGroup, nil
}

func (appGroup *appGroupTOML) constructConfigTree(tomlTree *toml.Tree, regExpAttr *regexp.Regexp, key string) error {
	// Get the top level keys for the tomlTree
	var topLevelKeys []string
	var isString, isNested bool = false, false
	var strVal string
	var intfVal interface{}
	var parentKey string
	var nestTOML *toml.Tree

	topLevelKeys = tomlTree.Keys()
	if len(topLevelKeys) <= 0 {
		// TOML file is blank. Something is wrong.
		return fmt.Errorf("Configuration File in toml is blank")
	}
	fmt.Println(appGroup.configTOMLTree)
	// Check for value substitution
	for _, topLevelKey := range topLevelKeys {
		// Is the topLevelKey a string
		intfVal = tomlTree.Get(topLevelKey)
		// fmt.Println(reflect.TypeOf(intfVal).String())
		if key != "" {
			parentKey = key + "." + topLevelKey
			fmt.Println(appGroup.targetTOMLTree.Has(parentKey))
			if !appGroup.targetTOMLTree.Has(parentKey) {
				appGroup.targetTOMLTree.Set(topLevelKey, intfVal)
			}
		} else {
			appGroup.targetTOMLTree.Set(topLevelKey, intfVal)
		}

		fmt.Println(appGroup.targetTOMLTree)
		fmt.Println(appGroup.configTOMLTree)
		nestTOML, isNested = intfVal.(*toml.Tree)
		if isNested {
			if parentKey != "" {
				appGroup.constructConfigTree(nestTOML, regExpAttr, parentKey)
			} else {
				appGroup.constructConfigTree(nestTOML, regExpAttr, topLevelKey)
			}
			continue
		}
		strVal, isString = tomlTree.Get(topLevelKey).(string)
		if !isString {
			arrTOML, isArrToml := intfVal.([]*toml.Tree)
			if isArrToml {
				if parentKey != "" {
					appGroup.constructConfigTree(arrTOML[0], regExpAttr, parentKey)
				} else {
					appGroup.constructConfigTree(arrTOML[0], regExpAttr, topLevelKey)
				}
			}
		}
		fmt.Println(appGroup.configTOMLTree)
		attrArray := regExpAttr.Split(strVal, -1)
		if len(attrArray) > 1 {
			arr := strings.Split(attrArray[1], ".")
			// This means that there is %developerInputs% match. i.e, this value needs to be substituted
			// Fetch the value from the developerInputs file
			// inputPath := attrArray[1] // Always in the 2nd subscript. Can simplify this
			inputPath := arr[0]
			inputArray, err := appGroup.getDeveloperInputs(inputPath)
			for i := 1; i < len(inputArray); i++ {
				configTOMLTree, ok := appGroup.configTOMLTree.Get(key).([]*toml.Tree)
				if !ok {
					return fmt.Errorf("Config TOML tree not []*toml.Tree")
				}
				targetTOMLTree, ok := appGroup.targetTOMLTree.Get(key).([]*toml.Tree)
				if !ok {
					return fmt.Errorf("Target TOML tree not []*toml.Tree")
				}
				appGroup.targetTOMLTree.Set(key, append(targetTOMLTree, configTOMLTree[0]))
			}

			fmt.Println(appGroup.targetTOMLTree)

			if err != nil {
				return fmt.Errorf("%v", err)
			}

			if len(inputArray) == 0 {
				return fmt.Errorf("%s", "No developer inputs present!")
			}
			// fmt.Println("TOML FILE : ", appGroup.targetTOMLTree, ";", appGroup.configTOMLTree)
			tree, isTree := appGroup.targetTOMLTree.Get(key).([]*toml.Tree)
			for index, inputVal := range inputArray {
				if inputVal != "" {
					if parentKey != "" {
						if isTree {
							tree[index].Set(parentKey, inputVal)
						} else {
							appGroup.targetTOMLTree.Set(parentKey, inputVal)
						}
					} else {
						appGroup.targetTOMLTree.Set(topLevelKey, inputVal)
					}
				}

			}

		}

	}

	return nil
}

func (appGroup *appGroupTOML) getDeveloperInputs(inputPath string) (inputArray []string, err error) {

	inputArray = []string{}
	var inputInf []interface{}
	defaultInputVal := appGroup.defaultTOMLTree.ToMap()[inputPath]

	fmt.Println(reflect.TypeOf(appGroup.developerTOMLTree.Get(inputPath)))

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

	} else if devInputStringInf, isStringInf := appGroup.developerTOMLTree.Get(inputPath).([]interface{}); isStringInf {
		if len(devInputStringInf) > 0 {
			for _, strVal := range devInputStringInf {
				fmt.Println(strVal)
				inputArray = append(inputArray, strVal.(string))
			}
		} else {
			if devInputStringInf, isDefaultStringInf := defaultInputVal.([]interface{}); isDefaultStringInf {
				inputInf = append(inputInf, devInputStringInf...)
			} else {
				return inputArray, fmt.Errorf("Values exist neither in input nor in default")
			}
		}

	} else {
		return inputArray, fmt.Errorf("Unable to fetch developer input string")
	}

	return inputArray, nil
}
