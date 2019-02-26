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


## IDE Error Logs
-2/25 & 26/19
```# command-line-arguments [command-line-arguments.test]
   ./kvcache_test.go:11:16: undefined: NewSimpleKVCache
   ./kvcache_test.go:23:17: undefined: SimpleKeyValueCache
   ./kvcache_test.go:39:17: undefined: SimpleKeyValueCache
   
   Compilation finished with exit code 2
```