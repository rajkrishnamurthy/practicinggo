package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	var err error

	dotStrings := []string{"1.0", "1.0.1", "1.1",
		"1.2", "1.3",
		"1.4", "1.4.1", "1.4.2",
		"1.4.2.2", "1.4.2.3",
		"1.5", "1.5.3", "1.5.3.5", "1.5.3.5.2",
	}

	depthMap := make(map[string][]string)

	for idx := 1; idx <= 6; idx++ {
		fmt.Println("----------")
		depthMap, err = modDotOutput(dotStrings, idx, depthMap)
		if err != nil {
			fmt.Printf("%v \n", err)
		}
	}
	fmt.Printf("Depth Map \n ------------ \n %v \n", depthMap)

}

func modDotOutput(inStrings []string, depth int, inMap map[string][]string) (outMap map[string][]string, err error) {

	weakDepth := strings.Repeat(`\.[0-9]`, func() int {
		if (depth - 2) < 0 {
			return 0
		} else {
			return depth - 2
		}
	}())
	weakPattern := `[0-9]\.[0-9]` + weakDepth + `+`
	rWeak, err := regexp.Compile(weakPattern)
	if err != nil {
		fmt.Println(err)
		return inMap, err
	}

	strongDepth := strings.Repeat(`\.\w{1}`, func() int {
		if (depth - 1) < 0 {
			return 0
		} else {
			return depth - 1
		}
	}())
	strongPattern := `^\w{1}\.\w{1}` + strongDepth + `$`
	rStrong, err := regexp.Compile(strongPattern)
	if err != nil {
		fmt.Println(err)
		return inMap, err
	}
	for _, eachStr := range inStrings {
		// fmt.Printf("Orig String = %v \t", eachStr)
		weakStringA := rWeak.FindAllString(eachStr, 1)
		weakString := func() string {
			if len(weakStringA) > 0 {
				return weakStringA[0]
			} else {
				return ""
			}
		}()
		// fmt.Printf("Weak String = %v \t \t", weakString)

		strongStringA := rStrong.FindAllString(eachStr, 1)
		strongString := func() string {
			if len(strongStringA) > 0 {
				return strongStringA[0]
			} else {
				return ""
			}
		}()

		// fmt.Printf("Strong String = %v \n", strongString)

		if strongString != "" {
			if len(inMap[weakString]) == 0 {
				inMap[weakString] = make([]string, 0)
			}
			inMap[weakString] = append(inMap[weakString], strongString)
		}

	}
	return inMap, nil
}
