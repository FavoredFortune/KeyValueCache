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

type simpleKeyValueCache struct{
	data map[string]string
}

//per discussion, thought about order of operations - needs to have a cache for the method to work
func (c *simpleKeyValueCache) Put(key,value string) error{
	if c !=nil {
		c.data[key]= value
		return nil
	}
	return fmt.Errorf("put failed: cache '%v' does not exists", c)
}

func (c *simpleKeyValueCache) Read(key string) string{
	return c.data[key]
}

func (c *simpleKeyValueCache) Update(key, value string) error{
	_, keyExists := c.data[key]
	if keyExists {
		c.data[key] = value
		return nil
	}
	return fmt.Errorf("update failed: key '%v' not in cache", key)
}

func (c *simpleKeyValueCache) Delete(key string) error{
	_, keyExist := c.data[key]
	if keyExist{
		delete(c.data, key)
		return nil
	}
	return fmt.Errorf("delete failed: key '%v' not in cache",key)
}

type toFileKeyValueCache struct{
	//members (to append reads to a file) - to check output
	//run till we get an error or check output
	//cache *simpleKeyValueCache - what ever it reads return it to file
}