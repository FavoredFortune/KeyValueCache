# KeyValueCache

## Purpose
- Work with Go
- Practice writing tests in Go
- Learn more about structs and interfaces
- Grow project over time to add other technologies and learning goals

## Technologies
- Go
- Goland (IDE)
- Stretchr/testify API (github.com/stretchr/testify)
- Cobra API (https://github.com/spf13/cobra)
- Flag API (from Go https://golang.org/pkg/flag/)
_more soon_

## User Stories
- As a user I want to put in code arguments like `put animal horse` that work with a Go CLI 'client' to take the commands in a terminal and add the values `animal` and `horse` as a `key:value` pair in this `SimpleKeyValueCache` construct
- As a user I want to type in the command `read animal` to the CLI and have the Go application return `>>horse`
- As a developer I want unit tests for each method (`Put`, `Read`, `Update`, `Delete`) that prove it works with both good and bad input
- As a developer I want to use the Cobra library to build the CLI (per Troy Dai)
* As a developer I want to use the Go Flag library in the CLI (per Troy & Scott)
- As a developer I want unit tests for each command 
- As a developer I want integrated tests for the CLI
- More developer stories details coming soon from Scott


## Expected command behaviors
- "put" puts a `key` string and a `value` string into a the designed simple value cache struct as a key:value pair
- "read" takes an input string `key` finds that in the cache (checking to be sure cache exists and key exists) and return it's paired `value` string
- "update" allows user to input any existing `key` string and change it's `value` string(after checking that cache and key exists)
- "delete" allows user to delete `key-value` string pair by inputting just the `key` from the cache (after checking that the cache and key exist)

## Future state of CLI
- Include flags for `key` that are `-k` or `-key` and `-v` and `-value` for value
- Have verbose manual option available with typing in command `man`


## Resources
- Mentors and co-workers: Scott Hornberger and Troy Dai
- gitignore generated on gitignore.io
- Article on testing in Go: https://medium.com/@matryer/5-simple-tips-and-tricks-for-writing-unit-tests-in-golang-619653f90742
- BETTER article on testing in Go with your own written unit tests: https://www.calhoun.io/how-to-test-with-go/
- On building a simple CLI in Go: https://blog.rapid7.com/2016/08/04/build-a-simple-cli-tool-with-golang/
- More on flags in Go: https://gobyexample.com/command-line-flags
- If I want to build a CLI using the Go CLI library: https://tutorialedge.net/golang/building-a-cli-in-go/

- For help on Cobra work around for auto generating command files - after using command with -t for package name, must go get file from gocode folder and put in app folder
https://github.com/spf13/cobra/pull/817

- Tutorial on building a CLI with Cobra: https://ordina-jworks.github.io/development/2018/10/20/make-your-own-cli-with-golang-and-cobra.html

- Make sub commands shared to convert to functions: https://stackoverflow.com/questions/43747075/cobra-commander-how-to-call-a-command-from-another-command

- Cobra Documentation: https://github.com/spf13/cobra

- Go Documentation on flags: https://golang.org/pkg/flag/

- More on flags: https://flaviocopes.com/go-command-line-flags/

#### Practice CLI project
- See https://github.com/FavoredFortune/CobraCLI for application with instructions in the [README](https://github.com/FavoredFortune/CobraCLI/blob/master/README.md)

## Reference code from earlier in the project

```gotemplate
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
```gotemplate
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

command?
```gotemplate
	//RunE: func(cmd *cobra.Command, args []string) error {
	//	switch args[0] {
	//
	//	case "put":
	//		if len(args) < 3 {
	//			return errors.New("put failed: put command and both key and value strings required")
	//		}
	//		cache.Put(args[1], args[2])
	//
	//	case "read":
	//		if len(args) < 2 {
	//			return fmt.Errorf("READ command incomplete: %v", args)
	//		}
	//		_, readResult := cache.Read(args[1])
	//		fmt.Println(">", readResult)
	//
	//	case "update":
	//		if len(args) < 3 {
	//			return fmt.Errorf("UPDATE command incomplete: %v", args)
	//		}
	//		err := cache.Update(args[1], args[2])
	//		if err != nil {
	//			return err
	//		}
	//
	//	case "delete":
	//		if len(args) < 2 {
	//			return fmt.Errorf("DELETE command incomplete: %v", args)
	//		}
	//		err := cache.Delete(args[1])
	//		if err != nil {
	//			return err
	//		}
	//	}
	//	return nil
	//},
```

