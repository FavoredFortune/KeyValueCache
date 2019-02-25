package kvcache

import (

	"testing"
)

func TestPut(t *testing.T) {
	cache := &simpleKeyValueCache{data: map[string]string{}}

	cache.Put("name is key","value is Troy Dai")

	if cache.Read("name is key") != "value is Troy Dai" {
		t.Fail()
	}
}