package main

import (
	"fmt"
)

func main() {
	//not sure if this is needed?
	testPut(handleInput(cache,testInputHappy), testOutputHappy)

}


var testInputHappy = `
name Sooz
`
var testOutputHappy = `
Sooz`

var cache = &simpleKeyValueCache{data: map[string]string{}}

func (c *simpleKeyValueCache) testPut( key,value, expected string) {

	//how do I call the Put method from the main file here?
	actual := Put (testInputHappy)

	if actual != testOutputHappy {
		message := fmt.Sprintf("Put failed: Got %v, but expected %v", actual, testOutputHappy)
		panic(message)
	}
}