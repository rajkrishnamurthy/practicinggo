package main

import (
	"fmt"
	"net/http"
	"reflect"
)

type httpreqobject int

func (f *httpreqobject) reflecttheobject(req *http.Request) {
	val := reflect.ValueOf(req).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag

		fmt.Printf("Field Name: %s,\t Field Value: %v,\t Tag Value: %s\n", typeField.Name, valueField.Interface(), tag.Get("tag_name"))
	}
}

func main() {
	http.HandleFunc("/", foofunc)
	http.ListenAndServe(":8080", nil)
}

func foofunc(w http.ResponseWriter, req *http.Request) {
	f := new(httpreqobject)
	// *f = req.(httpreqobject)
	f.reflecttheobject(req)
}
