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