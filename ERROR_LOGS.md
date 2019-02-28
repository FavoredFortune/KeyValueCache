## Terminal Error Logs
- 2/25/19
```
--- FAIL: TestPut (0.00s)
    --- FAIL: TestPut/it_can_put_and_read (0.00s)
        kvcache_test.go:30: 
            	Error Trace:	kvcache_test.go:30
            	Error:      	Received unexpected error:
            	            	put failed: key 'testKey' or value 'testValue' does not exist: 
            	Test:       	TestPut/it_can_put_and_read
--- FAIL: TestRead (0.00s)
    --- FAIL: TestRead/it_can_read (0.00s)
        kvcache_test.go:47: 
            	Error Trace:	kvcache_test.go:47
            	Error:      	Received unexpected error:
            	            	put failed: key 'name' or value 'Scott' does not exist: 
            	Test:       	TestRead/it_can_read
        kvcache_test.go:49: 
            	Error Trace:	kvcache_test.go:49
            	Error:      	Not equal: 
            	            	expected: ""
            	            	actual  : "Scott"
            	            	
            	            	Diff:
            	            	--- Expected
            	            	+++ Actual
            	            	@@ -1 +1 @@
            	            	-
            	            	+Scott
            	Test:       	TestRead/it_can_read
FAIL
exit status 1
FAIL	_/Users/srichm/Uber/KVCache/kvcache	0.016s
```
-2/26/19 - 10 am
```=== RUN   TestSimpleKeyValueCache
   === RUN   TestSimpleKeyValueCache/new_cache_created
   --- PASS: TestSimpleKeyValueCache (0.00s)
       --- PASS: TestSimpleKeyValueCache/new_cache_created (0.00s)
   === RUN   TestPut
   === RUN   TestPut/it_can_put_and_read
   === RUN   TestPut/second_put_test
   --- PASS: TestPut (0.00s)
       --- PASS: TestPut/it_can_put_and_read (0.00s)
       --- PASS: TestPut/second_put_test (0.00s)
   === RUN   TestRead
   === RUN   TestRead/it_can_read
   === RUN   TestRead/second_read_test
   --- FAIL: TestRead (0.00s)
       --- PASS: TestRead/it_can_read (0.00s)
       --- FAIL: TestRead/second_read_test (0.00s)
           kvcache_test.go:83: 
               	Error Trace:	kvcache_test.go:83
               	Error:      	Not equal: 
               	            	expected: *errors.errorString(&errors.errorString{s:"read failed: key '' invalid"})
               	            	actual  : string("read failed: key ' ' invalid")
               	Test:       	TestRead/second_read_test
   FAIL
   exit status 1
   FAIL	_/Users/srichm/Uber/KVCache/kvcache	0.016s
   ```
-2/26/10 10:38am (testing Update)
```=== RUN   TestSimpleKeyValueCache
   === RUN   TestSimpleKeyValueCache/new_cache_created
   --- PASS: TestSimpleKeyValueCache (0.00s)
       --- PASS: TestSimpleKeyValueCache/new_cache_created (0.00s)
   === RUN   TestPut
   === RUN   TestPut/it_can_put_and_read
   === RUN   TestPut/second_put_test
   --- PASS: TestPut (0.00s)
       --- PASS: TestPut/it_can_put_and_read (0.00s)
       --- PASS: TestPut/second_put_test (0.00s)
   === RUN   TestRead
   === RUN   TestRead/it_can_read
   === RUN   TestRead/second_read_test
   --- PASS: TestRead (0.00s)
       --- PASS: TestRead/it_can_read (0.00s)
       --- PASS: TestRead/second_read_test (0.00s)
   === RUN   TestUpdate
   === RUN   TestUpdate/it_can_update
   --- FAIL: TestUpdate (0.00s)
       --- FAIL: TestUpdate/it_can_update (0.00s)
           kvcache_test.go:107: 
               	Error Trace:	kvcache_test.go:107
               	Error:      	Not equal: 
               	            	expected: <nil>(<nil>)
               	            	actual  : string("Benny")
               	Test:       	TestUpdate/it_can_update
   FAIL
   exit status 1
   FAIL	_/Users/srichm/Uber/KVCache/kvcache	0.014s
   ```
-2/26/19 10:50 (3rd Update test - empty key)
```=== RUN   TestSimpleKeyValueCache
   === RUN   TestSimpleKeyValueCache/new_cache_created
   --- PASS: TestSimpleKeyValueCache (0.00s)
       --- PASS: TestSimpleKeyValueCache/new_cache_created (0.00s)
   === RUN   TestPut
   === RUN   TestPut/it_can_put_and_read
   === RUN   TestPut/second_put_test
   --- PASS: TestPut (0.00s)
       --- PASS: TestPut/it_can_put_and_read (0.00s)
       --- PASS: TestPut/second_put_test (0.00s)
   === RUN   TestRead
   === RUN   TestRead/it_can_read
   === RUN   TestRead/second_read_test
   --- PASS: TestRead (0.00s)
       --- PASS: TestRead/it_can_read (0.00s)
       --- PASS: TestRead/second_read_test (0.00s)
   === RUN   TestUpdate
   === RUN   TestUpdate/it_can_update
   === RUN   TestUpdate/update_error_works
   === RUN   TestUpdate/empty_key_Update_test
   --- FAIL: TestUpdate (0.00s)
       --- PASS: TestUpdate/it_can_update (0.00s)
       --- PASS: TestUpdate/update_error_works (0.00s)
       --- FAIL: TestUpdate/empty_key_Update_test (0.00s)
           kvcache_test.go:138: 
               	Error Trace:	kvcache_test.go:138
               	Error:      	Not equal: 
               	            	expected: *errors.errorString(&errors.errorString{s:"update failed: key '' not in cache"})
               	            	actual  : <nil>(<nil>)
               	Test:       	TestUpdate/empty_key_Update_test
   FAIL
   exit status 1
   ```
- 2/27/19 2:37pm
Trying to put server/handler logic in root command and it creates a panic...likely because sometimes the args isn't log enough for some if statements in the switch cases?
```gotemplate
go run main.go
panic: runtime error: index out of range

goroutine 1 [running]:
KVCache/cmd.glob..func1(0x17eab60, 0x1810360, 0x0, 0x0, 0x0, 0x0)
	/Users/srichm/gocode/src/KVCache/cmd/root.go:40 +0x5af
github.com/spf13/cobra.(*Command).execute(0x17eab60, 0xc00008c190, 0x0, 0x0, 0x17eab60, 0xc00008c190)
	/Users/srichm/gocode/src/github.com/spf13/cobra/command.go:762 +0x473
github.com/spf13/cobra.(*Command).ExecuteC(0x17eab60, 0x1495bf0, 0x1, 0x0)
	/Users/srichm/gocode/src/github.com/spf13/cobra/command.go:852 +0x2fd
github.com/spf13/cobra.(*Command).Execute(0x17eab60, 0x1007220, 0xc000082058)
	/Users/srichm/gocode/src/github.com/spf13/cobra/command.go:800 +0x2b
KVCache/cmd.Execute()
	/Users/srichm/gocode/src/KVCache/cmd/root.go:94 +0x2d
main.main()
	/Users/srichm/gocode/src/KVCache/main.go:22 +0x20
exit status 2
```
- 2/27/19 2:51 PM Read command appears to work, but proves put command doesn't work (error does, but it doesn't add content to KVC)
```gotemplate
srichm :~/gocode/src/KVCache :[weds-cli] go run main.go
This is a CLI app that allows you to input your action and your key;value pair of strings.

	Actions available include put, read,  update or delete your content from the cache.

srichm :~/gocode/src/KVCache :[weds-cli !] ./main put horse animal
srichm :~/gocode/src/KVCache :[weds-cli !] ./main read horse
> read failed: key 'horse' invalid
srichm :~/gocode/src/KVCache :[weds-cli !] ./main read animal
> read failed: key 'animal' invalid
srichm :~/gocode/src/KVCache :[weds-cli !] 
```

- 2/27/19 4:13 PM Error message in put working with simplified code but put command itself not putting anything. 
```Error: >>put failed: put command and both key and value strings required
   Usage:
     cli put [flags]
   
   Flags:
     -h, --help   help for put
   
   srichm :~/gocode/src/KVCache :[weds-cli !x?] cli put name steve
   -bash: cli: command not found
   srichm :~/gocode/src/KVCache :[weds-cli !x?] ./cli put
   Error: >>put failed: put command and both key and value strings required
   Usage:
     cli put [flags]
   
   Flags:
     -h, --help   help for put
   
   srichm :~/gocode/src/KVCache :[weds-cli !x?] ./cli put animal horse
   Error: >>put failed: put command and both key and value strings required
   Usage:
     cli put [flags]
   
   Flags:
     -h, --help   help for put
```

- 2/27/19 4:47pm
Simple flags operating on their own, but not incorporated to put correctly - YET!
```srichm :~/gocode/src/KVCache :[weds-cli !] ./cli put name bene
  Error: >>put failed: put command and both key and value strings required
  Usage:
    cli put [flags]
  
  Flags:
    -h, --help   help for put
  
  cmd: args
  tail: [put name bene]
```
- 2/28/19 6:00 PM
Flag for put is starting to work, but not parsing args as expected, losing second part of string slice
```srichm :~/gocode/src/KVCache :[weds-cli !] ./cli -put name bene
   Error: unknown command "name" for "cli"
   Run 'cli --help' for usage.
   put: name
   srichm :~/gocode/src/KVCache :[weds-cli !] ./cli
   Usage:
     cli [command]
   
   Available Commands:
     help        Help about any command
     put         put key-value pair
     read        read value given key
   
   Flags:
     -h, --help   help for cli
   
   Use "cli [command] --help" for more information about a command.
   put: 
   srichm :~/gocode/src/KVCache :[weds-cli !] ./cli --put name bene
   Error: unknown command "bene" for "cli"
   Run 'cli --help' for usage.
   put: name
   srichm :~/gocode/src/KVCache :[weds-cli !] ./cli name bene --put
   Error: unknown command "name" for "cli"
   Run 'cli --help' for usage.
   put: 
   srichm :~/gocode/src/KVCache :[weds-cli !] ./cli name --put bene
   Error: unknown command "name" for "cli"
   Run 'cli --help' for usage.
   put: 

```

- 2/27/19 6:09 PM flag.Parse and flag.Args appear to work correctly if default value is string with length of 2, but still doesn't work with command of length of two.
```srichm :~/gocode/src/KVCache :[weds-cli !] ./cli --put name bene
   Error: unknown command "bene" for "cli"
   Run 'cli --help' for usage.
   put: name
   srichm :~/gocode/src/KVCache :[weds-cli !] ./cli --put
   Error: unknown flag: --put
   Usage:
     cli [command]
   
   Available Commands:
     help        Help about any command
     put         put key-value pair
     read        read value given key
   
   Flags:
     -h, --help   help for cli
   
   Use "cli [command] --help" for more information about a command.
   
   flag needs an argument: -put
   Usage of ./cli:
     -put string
       	a key-value pair of strings to put in cache (Required) (default "key value")

```
- 2/27/19 6:16 PM
added init func and put flags there to see if it helped with parsing the args, but it doesn't change functionality..at least it doesn't break anything. 

```
srichm :~/gocode/src/KVCache :[weds-cli !] go build -o cli
   srichm :~/gocode/src/KVCache :[weds-cli !] ./cli --put
   flag needs an argument: -put
   Usage of ./cli:
     -put string
       	a key-value pair of strings to put in cache (Required) (default "key value")
   srichm :~/gocode/src/KVCache :[weds-cli !] ./cli --put name bene
   put: name
   Error: unknown command "bene" for "cli"
   Run 'cli --help' for usage.
   srichm :~/gocode/src/KVCache :[weds-cli !] 
   ```
   
-2/28/19 10:46 AM
Put command working as expected - can even call read method within it - but Read command erroring out
```srichm :~/gocode/src/KVCache :[thurs-cli !] go build -o cli
   srichm :~/gocode/src/KVCache :[thurs-cli !] ./cli put name Bene
   [name Bene] 2
    key 'name' and value 'Bene' put into the cache
   Bene <nil>
   srichm :~/gocode/src/KVCache :[thurs-cli !] ./cli read name
   [name] 1
   Error: read failed: key 'name' invalid or cache empty
   Usage:
     cli read [flags]
   
   Flags:
     -h, --help   help for read
 ```

## IDE Error Logs
- 2/25 & 26/19
```# command-line-arguments [command-line-arguments.test]
   ./kvcache_test.go:11:16: undefined: NewSimpleKVCache
   ./kvcache_test.go:23:17: undefined: SimpleKeyValueCache
   ./kvcache_test.go:39:17: undefined: SimpleKeyValueCache
   
   Compilation finished with exit code 2
```
Above error is happening due to IDE looking for root and path in wrong folders per Troy Dai.
- 2/27/19
   * Fixed IDE error by changing project structure and go env variables
   * IDE tests work now


