package main

func main() {
	testPut(handleInput(cache,testInputHappy), testOutputHappy)

}


var testInputHappy = `
name Sooz
`
var testOutputHappy = `
Sooz`

var cache = &simpleKeyValueCache{data: map[string]string{}}

func (c *simpleKeyValueCache) testPut( key,value, expected string) error {
	actual := Put(testInputHappy)

	if actual !
}