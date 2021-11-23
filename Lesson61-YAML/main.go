package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v3"
)

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
// type T struct {
// 	A string
// 	B struct {
// 		RenamedC int   `yaml:"c"`
// 		D        []int `yaml:",flow"`
// 	}
// }

func main() {
	var err error
	var d []byte

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal([]byte(YAMLInput), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m:\n%v\n\n", m)

	spec, _ := m["spec"]
	template, _ := spec.(map[string]interface{})["template"]
	metadata, _ := template.(map[string]interface{})["metadata"]
	labels, _ := metadata.(map[string]interface{})["labels"].(map[string]interface{})

	for k, v := range labels {
		fmt.Printf("spec.template.metadata.labels: Key = %s\tValue=%s\n", k, v)
	}

	d, err = yaml.Marshal(&m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m dump:\n%s\n\n", string(d))
}
