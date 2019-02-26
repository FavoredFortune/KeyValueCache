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