package kvcache

import (
	"fmt"
)

//interface for use by all files (public by using cap at start of name)
type KeyValueCache interface{
	Put(key, value string) error
	Read(key string) (string,error)
	Update(key,value string) error
	Delete(key string) error
}

type SimpleKeyValueCache struct{
	data map[string]string
}

//Leaving in constructor function in case it's useful later
func NewSimpleKVCache() *SimpleKeyValueCache{
	return &SimpleKeyValueCache{map[string]string{}}
}

//per Troy don't need to check for cache here, this is a method of c - it is like "'this'in Java"
func (c *SimpleKeyValueCache) Put(key,value string) error{

	//added if statement to match read behavior and logic for empty string
	if key =="" || value =="" {
		return fmt.Errorf("put failed: check key '%v' and value '%v' parameters  ",key, value)
	}
	_, keyExists := c.data[key]
	if keyExists {
		return fmt.Errorf("put failed: key '%v' isn't unqiue: ", key)
	}
	c.data[key]=value
	return nil
}

//updated interface and method to return both string and error when realized SKVC wouldn't return an error when an empty string was entered as a key - not cool
func (c *SimpleKeyValueCache) Read(key string) (string,error){
	f, err := c.data[key]
	if err != true {
		return "",fmt.Errorf("read failed: key '%v' invalid", key)
	}
	return f, nil
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
