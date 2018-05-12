package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type ProductionInfo struct {
	StructA []Entry
}

type Entry struct {
	Field1 string
	Field2 int
	Field3 []string
	Field4 []byte
}

type CNIOType string

func PrintEntry(entry Entry) {
	fmt.Printf("Inside the Print Entry Function \n")
	fmt.Printf("%v \n", entry)
}

func initSource(source interface{}) {
	var refSourceVal, refFieldVal reflect.Value

	refSourceVal = reflect.ValueOf(source)
	fmt.Printf("Type = %v \t Kind = %v \n", refSourceVal.Type().String(), refSourceVal.Kind().String())
	if refSourceVal.Kind() == reflect.Ptr {
		refSourceVal = refSourceVal.Elem()
	}
	fmt.Printf("Type = %v \t Kind = %v \n", refSourceVal.Type().String(), refSourceVal.Kind().String())
	if refSourceVal.Kind() != reflect.Struct {
		return
	}

	for idx := 0; idx < refSourceVal.NumField(); idx++ {
		fmt.Printf("Type Name = %s \t Field Name = %s \n", refSourceVal.Type().Name(), refSourceVal.Type().Field(idx).Name)
		refFieldVal = refSourceVal.Field(idx)
		fmt.Printf("refFieldVal \t Name = %v \t Type =  %v \t Kind = %v \t Settable = %v \n", refFieldVal.Type().Name(), refFieldVal.Type().String(), refFieldVal.Kind().String(), refFieldVal.CanSet())
		if refFieldVal.CanSet() {
			switch refFieldVal.Kind() {
			case reflect.String, reflect.Int:
				refFieldVal.Set(reflect.Zero(refFieldVal.Type()))
			case reflect.Slice:
				refFieldVal.Set(reflect.MakeSlice(refFieldVal.Type(), 1, 1))
				// refFieldVal.Set(reflect.Zero(reflect.SliceOf(refFieldVal.Type())))

			default:
			}

		}
	}

	PrintEntry(refSourceVal.Interface().(Entry))
}

func SetField(source interface{}) {

	var refSourceVal, refFieldVal reflect.Value

	refSourceVal = reflect.ValueOf(source)
	fmt.Printf("Type = %v \t Kind = %v \n", refSourceVal.Type().String(), refSourceVal.Kind().String())
	if refSourceVal.Kind() == reflect.Ptr {
		refSourceVal = refSourceVal.Elem()
	}
	fmt.Printf("Type = %v \t Kind = %v \n", refSourceVal.Type().String(), refSourceVal.Kind().String())
	if refSourceVal.Kind() != reflect.Struct {
		return
	}

	for idx := 0; idx < refSourceVal.NumField(); idx++ {
		fmt.Printf("Type Name = %s \t Field Name = %s \t Field Type = %s \n", refSourceVal.Type().Name(), refSourceVal.Type().Field(idx).Name, refSourceVal.Type().Field(idx).Type)
		refFieldVal = refSourceVal.Field(idx)
		fmt.Printf("refFieldVal \t Name = %v \t Type =  %v \t Kind = %v \t Settable = %v \n", refFieldVal.Type().Name(), refFieldVal.Type().String(), refFieldVal.Kind().String(), refFieldVal.CanSet())
		if refFieldVal.CanSet() {
			switch refFieldVal.Kind() {
			case reflect.String:
				refFieldVal.SetString("SomeValue")
			case reflect.Int:
				refFieldVal.SetInt(122)
			case reflect.Slice:

				fmt.Printf("Slice.Type.Elem.Kind = %v  \n",
					refFieldVal.Type().Elem().Kind().String(),
				)
				// refFieldVal.Set(reflect.MakeSlice(refFieldVal.Type(), 5, 5))
				switch refFieldVal.Type().Elem().Kind() {
				case reflect.String:
					refFieldVal.Set(reflect.Append(refFieldVal, reflect.ValueOf("entry"+strconv.Itoa(idx))))
					refFieldVal.Set(reflect.Append(refFieldVal, reflect.ValueOf("exit"+strconv.Itoa(idx+2))))

				case reflect.Uint8:
					var test uint8 = 10
					refFieldVal.Set(reflect.Append(refFieldVal, reflect.ValueOf(test)))
					refFieldVal.Set(reflect.Append(refFieldVal, reflect.ValueOf(test)))
				default:
				}

			default:
			}

		}
	}

	PrintEntry(refSourceVal.Interface().(Entry))
}

func main() {
	// source := ProductionInfo{}
	// source.StructA = append(source.StructA, Entry{Field1: "A", Field2: 2})

	// fmt.Println("Before: ", source.StructA[0])
	// SetField(&source.StructA[0])
	// fmt.Println("After: ", source.StructA[0])

	source := &Entry{}
	// SetField(source)
	initSource(source)
	fmt.Println(" :", source)
}
