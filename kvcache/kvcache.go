package kvcache

import (
	"fmt"
)

//interface for use by all files (public by using cap at start of name)
type KeyValueCache interface{
	Put(key, value string) error
	Read(key string) string
	Update(key,value string) error
	Delete(key string) error
}

type SimpleKeyValueCache struct{
	data map[string]string
}

//AHA! I was missing the constructor function for the cache! This is why tests were breaking (I hope)
func NewSimpleKVCache() *SimpleKeyValueCache{
	return &SimpleKeyValueCache{map[string]string{}}
}

//per Troy don't need to check for cache here, this is a method of c - it is like "this in Java"
func (c *SimpleKeyValueCache) Put(key,value string) error{

			_, keyExists := c.data[key]
			if keyExists {
				c.data[key] = value
				return nil
			}
		return fmt.Errorf("put failed: key '%v' or value '%v' does not exist: ",key, value)
}

func (c *SimpleKeyValueCache) Read(key string) string{
	return c.data[key]
}

func (c *SimpleKeyValueCache) Update(key, value string) error{
	_, keyExists := c.data[key]
	if keyExists {
		c.data[key] = value
		return nil
	}
	return fmt.Errorf("update failed: key '%v' not in cache", key)
}

func (c *SimpleKeyValueCache) Delete(key string) error{
	_, keyExist := c.data[key]
	if keyExist{
		delete(c.data, key)
		return nil
	}
	return fmt.Errorf("delete failed: key '%v' not in cache",key)
}
//adding comments to get changes into git