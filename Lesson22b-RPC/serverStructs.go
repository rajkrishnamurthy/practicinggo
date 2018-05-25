package main

const (
	_ = iota
	schemaInputs
	schemaFacts
	schemaOutputs
)

type InputValues struct {
	Inputs CNTaskVarSchema
	Facts  CNTaskVarSchema
	Output CNTaskVarSchema
}

// type OutputValues struct {
// 	Output CNTaskVarSchema
// }

type CNTaskVarSchema struct { // This is derived from cnTaskIOSchema for inputs
	Names []CNTaskIOName // contains an array of input variable names
	// Descriptions map[CNTaskIOName]string       // contains a key-value map of input types
	// Types        map[CNTaskIOName]CNTaskIOType // contains a key-value map of input types
	// Values       map[CNTaskIOName]string       // contains the default value for the input variable
	// Output       map[CNTaskIOName][]byte       // Seperating Output from Values
	Code []byte // contains the respective struct code
}

type OutputValues struct {
	// Names []string // contains an array of input variable names
	// Code  []byte   // contains the respective struct code
	Name string
}

type CNTaskIOType string // Types of Task Inputs/Outputs
type CNTaskIOName string

type TaskInstance int
