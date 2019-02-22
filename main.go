package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	testInput := `
PUT name Sooz
READ name
DELETE name
READ name
UPDATE name Suzanne
READ name
`
	cache := &simpleKeyValueCache{data: map[string]string{}}

	err := handleInput(cache, testInput)
	if err != nil {
		fmt.Println(err)
	}


}

func handleInput(cache KeyValueCache, input string) error {
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		cmd := strings.Split(scanner.Text(), " ")
		switch cmd[0] {

		case "PUT":
			if len(cmd) < 3{
				return fmt.Errorf("PUT command incomplete: %v", cmd)
			}
			err := cache.Put(cmd[1], cmd[2])
			if err != nil {
				return err
			}

		case "READ":
			if len(cmd) < 2{
				return fmt.Errorf("READ command incomplete: %v", cmd)
			}

			readResult := cache.Read(cmd[1])
			fmt.Println(">", readResult)

		case "UPDATE":
			if len(cmd) < 3{
				return fmt.Errorf("UPDATE command incomplete: %v", cmd)
			}
			err := cache.Update(cmd[1],cmd[2])
			if err != nil {
				return err
			}


		case "DELETE":
			if len(cmd) < 2 {
				return fmt.Errorf("DELETE command incomplete: %v", cmd)
			}
			err := cache.Delete(cmd[1])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

type KeyValueCache interface{
	Put(key, value string) error
	Read(key string) string
	Update(key,value string) error
	Delete(key string) error
}

type simpleKeyValueCache struct{
	data map[string]string
}


func (c *simpleKeyValueCache) Put(key,value string) error{
	c.data[key] = value
	return nil
}

func (c *simpleKeyValueCache) Read(key string) (string){
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