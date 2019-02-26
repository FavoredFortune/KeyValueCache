# KeyValueCache

## Purpose
- Work with Go
- Practice writing tests in Go
- Learn more about structs and interfaces
- Grow project over time to add other technologies and learning goals

## Technologies
- Go
- Goland (IDE)
- _more soon_

## User Stories

- Developer story details coming soon from Scott




## Resources
- Mentors and co-workers: Scott Hornberger and Troy Dai
- gitignore generated on gitignore.io
- Article on testing in Go: https://medium.com/@matryer/5-simple-tips-and-tricks-for-writing-unit-tests-in-golang-619653f90742
- BETTER article on testing in Go with your own written unit tests: https://www.calhoun.io/how-to-test-with-go/


## Reference code from earlier in the project

```go
func kvcache() {
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
```
```go
func handleInput(cache KeyValueCache, input string) error {
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		cmd := strings.Split(scanner.Text(), " ")
		switch cmd[0] {

		case "PUT":
			if len(cmd) < 3{
				return fmt.Errorf("PUT command incomplete: %v", cmd)
			}
			cache.Put(cmd[1], cmd[2])

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
```

