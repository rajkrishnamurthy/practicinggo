package main

// Obj1 : Used for Command Execution
type Obj1 struct {
	Name  string
	ID    int
	Input struct {
		Cmd      string
		Params   []string
		callfunc *func(*Obj1, *Obj2) bool
	}
}

// Obj2 : Input/Output Model
type Obj2 struct {
	Output struct {
		OutString string
		OutBytes  []byte
	}
}

// Taskinstance : Global receiver
type Taskinstance struct {
	InstanceName string
	InstanceType string
}
