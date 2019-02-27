##Terminal Test Logs
-2/26/19 9:18 am
```
=== RUN   TestSimpleKeyValueCache
=== RUN   TestSimpleKeyValueCache/new_cache_created
--- PASS: TestSimpleKeyValueCache (0.00s)
    --- PASS: TestSimpleKeyValueCache/new_cache_created (0.00s)
=== RUN   TestPut
=== RUN   TestPut/it_can_put_and_read
--- PASS: TestPut (0.00s)
    --- PASS: TestPut/it_can_put_and_read (0.00s)
=== RUN   TestRead
=== RUN   TestRead/it_can_read
--- PASS: TestRead (0.00s)
    --- PASS: TestRead/it_can_read (0.00s)
PASS
ok  	_/Users/srichm/Uber/KVCache/kvcache	0.016s

```

-2/26/19 10 am
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
   PASS
   ok  	_/Users/srichm/Uber/KVCache/kvcache	0.014s
   ```
- 2/26/19 10:40 am (1st Update test passes)
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
   --- PASS: TestUpdate (0.00s)
       --- PASS: TestUpdate/it_can_update (0.00s)
   PASS
   ok  	_/Users/srichm/Uber/KVCache/kvcache	0.017s
   
```
-2/26/19 (2nd update test) 10:47 am
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
   === RUN   TestUpdate/Update_error_works
   --- PASS: TestUpdate (0.00s)
       --- PASS: TestUpdate/it_can_update (0.00s)
       --- PASS: TestUpdate/Update_error_works (0.00s)
   PASS
   ok  	_/Users/srichm/Uber/KVCache/kvcache	0.019s
   ```
-2/26/19 10:51 AM (3rd Update test for error working with empty key input)
```
=== RUN   TestSimpleKeyValueCache
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
=== RUN   TestUpdate/empty_key_Update_error_test
--- PASS: TestUpdate (0.00s)
    --- PASS: TestUpdate/it_can_update (0.00s)
    --- PASS: TestUpdate/update_error_works (0.00s)
    --- PASS: TestUpdate/empty_key_Update_error_test (0.00s)
PASS
ok  	_/Users/srichm/Uber/KVCache/kvcache	0.017s
```
2/26/19 11:04 AM (added Put tests after logic of KVC re-evaluation per changes to Read method - i.e. failure with empty strings - and subsequent changes to Put method)
```=== RUN   TestSimpleKeyValueCache
   === RUN   TestSimpleKeyValueCache/new_cache_created
   --- PASS: TestSimpleKeyValueCache (0.00s)
       --- PASS: TestSimpleKeyValueCache/new_cache_created (0.00s)
   === RUN   TestPut
   === RUN   TestPut/it_can_put_and_read
   === RUN   TestPut/second_put_test
   === RUN   TestPut/_put_test_for_error_working
   --- PASS: TestPut (0.00s)
       --- PASS: TestPut/it_can_put_and_read (0.00s)
       --- PASS: TestPut/second_put_test (0.00s)
       --- PASS: TestPut/_put_test_for_error_working (0.00s)
   === RUN   TestRead
   === RUN   TestRead/it_can_read
   === RUN   TestRead/read_test_for_error_working
   --- PASS: TestRead (0.00s)
       --- PASS: TestRead/it_can_read (0.00s)
       --- PASS: TestRead/read_test_for_error_working (0.00s)
   === RUN   TestUpdate
   === RUN   TestUpdate/it_can_update
   === RUN   TestUpdate/update_error_works
   === RUN   TestUpdate/empty_key_Update_error_test
   --- PASS: TestUpdate (0.00s)
       --- PASS: TestUpdate/it_can_update (0.00s)
       --- PASS: TestUpdate/update_error_works (0.00s)
       --- PASS: TestUpdate/empty_key_Update_error_test (0.00s)
   PASS
   ok  	_/Users/srichm/Uber/KVCache/kvcache	0.016s
   ```
- 2/26/19 1:14PM (all delete unit tests added and passing)

```
=== RUN   TestSimpleKeyValueCache
   === RUN   TestSimpleKeyValueCache/new_cache_created
   --- PASS: TestSimpleKeyValueCache (0.00s)
       --- PASS: TestSimpleKeyValueCache/new_cache_created (0.00s)
   === RUN   TestPut
   === RUN   TestPut/it_can_put_and_read
   === RUN   TestPut/second_put_test
   === RUN   TestPut/_put_test_for_error_working
   --- PASS: TestPut (0.00s)
       --- PASS: TestPut/it_can_put_and_read (0.00s)
       --- PASS: TestPut/second_put_test (0.00s)
       --- PASS: TestPut/_put_test_for_error_working (0.00s)
   === RUN   TestRead
   === RUN   TestRead/it_can_read
   === RUN   TestRead/read_test_for_diff_keys
   === RUN   TestRead/read_test_for_error_working
   --- PASS: TestRead (0.00s)
       --- PASS: TestRead/it_can_read (0.00s)
       --- PASS: TestRead/read_test_for_diff_keys (0.00s)
       --- PASS: TestRead/read_test_for_error_working (0.00s)
   === RUN   TestUpdate
   === RUN   TestUpdate/it_can_update
   === RUN   TestUpdate/update_error_works
   === RUN   TestUpdate/empty_key_Update_error_test
   --- PASS: TestUpdate (0.00s)
       --- PASS: TestUpdate/it_can_update (0.00s)
       --- PASS: TestUpdate/update_error_works (0.00s)
       --- PASS: TestUpdate/empty_key_Update_error_test (0.00s)
   === RUN   TestDelete
   === RUN   TestDelete/it_deletes
   === RUN   TestDelete/delete_error_test
   === RUN   TestDelete/delete_error_test_with_empty_key_string
   --- PASS: TestDelete (0.00s)
       --- PASS: TestDelete/it_deletes (0.00s)
       --- PASS: TestDelete/delete_error_test (0.00s)
       --- PASS: TestDelete/delete_error_test_with_empty_key_string (0.00s)
   ```
   
- 2/27/19 2:44 PM CLI terminal test of command "put" (not quite right but close)
```
go run main.go
This is a CLI app that allows you to input your action and your key;value pair of strings.

	Actions available include put, read,  update or delete your content from the cache.

srichm :~/gocode/src/KVCache :[weds-cli !] ./main put horse animal
srichm :~/gocode/src/KVCache :[weds-cli !] ./main put
Error: put failed: put command and both key and value strings required
Usage:
  cli [flags]

Flags:
      --config string   config file (default is $HOME/.KVCache.yaml)
  -h, --help            help for cli
  -t, --toggle          Help message for toggle

put failed: put command and both key and value strings required
```