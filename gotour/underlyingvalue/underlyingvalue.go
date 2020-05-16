package main

import (
	"fmt"
)

func assert(i interface{}) {
	v, ok := i.(string) //get the underlying int value from i
	fmt.Println(v, ok)
}
func main() {
	var s interface{} = 56
	assert(s)
	var i interface{} = "Steven Paul"
	assert(i)
}
