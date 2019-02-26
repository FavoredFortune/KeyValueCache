package kvcache

import (
	"fmt"
	"testing"
)

func TestPut(t *testing.T) {
	//IDE testing isn't recognizing cache, but command line tests are running correctly
	cache := &SimpleKeyValueCache{ map[string]string{}}

	if cache ==nil {
		fmt.Println("cache is broken")
	}

	//testing long strings
	if error(cache.Put("name is key","value is Troy Dai")) !=nil {
		t.Errorf("Put Test failed: %v, %v, %v", cache)
	}

	////testing one empty string
	//cache.Put("","Steve")
	//if cache.Read("") != "Steve"{
	//	t.Fail()
	//}
	////testing one empty string
	//cache.Put("name","")
	//if cache.Read("name") != ""{
	//	t.Fail()
	//}

}