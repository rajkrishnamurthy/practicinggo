package main

import (
	"fmt"
	"sort"
)

type people sort.StringSlice

func (mypeeps people) Len() int {
	return len(mypeeps)
}

func (mypeeps people) Less(i, j int) bool {
	return (mypeeps[i] < mypeeps[j])
}

func (mypeeps people) Swap(i, j int) {
	var temp string
	temp = mypeeps[i]
	mypeeps[i] = mypeeps[j]
	mypeeps[j] = temp
}

func main() {

	// s := []string{"raj", "amigo", "loga", "shradha", "shruti", "subha"}
	// fmt.Println(sort.StringsAreSorted(s))
	// sort.Strings(s)
	// fmt.Println(s)

	mypeople := people{"raj", "amigo", "loga", "shradha", "shruti", "subha"}
	sort.Sort(mypeople)
	fmt.Println(mypeople)

}
