package main

import (
	"code.uber.internal/sooz/key-value/.tmp/.go/goroot/src/testing"
	"fmt"
)

func main() {
	
}

type TestKeyValueCache interface{
	Put(key, value string) error
	Read(key string) string
	Update(key,value string) error
	Delete(key string) error
}

type TestSimpleKeyValueCache struct{
	data map[string]string
}

var _ package.TestSimpleKeyValueCache = (*testTestSimpleKeyValueCache)(nil)

var handleInputKVCache = &TestSimpleKeyValueCache{ map[string]string{}}

var testInputHappy = `
PUT name Sooz
READ name
UPDATE name Betty
READ name
DELETE name
`
var testOutputHappy = `
Sooz
Betty`

func (c *TestSimpleKeyValueCache) TestHandleInput(t *testing.T)error {
	//create actual value using function
	actual := handleInput(handleInputKVCache, testInputHappy)

	if actual.Error() != nil{
		t.Errorf("Error: handleInput function not working. Expected % v, actual %v", actual, testOutputHappy)
	}

	fmt.Println("SUCCESS")

	return nil



}