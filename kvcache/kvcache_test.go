package kvcache

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSimpleKeyValueCache(t *testing.T) {
	t.Run("new cache created", func(t *testing.T) {
		testCache := NewSimpleKVCache()
		assert.NotNil(t, testCache)
	})

}


//not working going to try writing other tests and revisit
func TestPut(t *testing.T) {

	t.Run("it can put and read", func(t *testing.T) {

		testCache := &SimpleKeyValueCache{map[string]string{}}
		require.NotNil(t, testCache)

		key := "testKey"
		value := "testValue"
		err := testCache.Put(key,value)

		assert.NoError(t,err)
		assert.ObjectsAreEqualValues(testCache.Read(key),value)

	})
}

//also doesn't work
func TestRead(t *testing.T){
	t.Run("it can read", func(t *testing.T) {
		testCache := &SimpleKeyValueCache{map[string]string{}}
		require.NotNil(t, testCache)

		key := "name"
		value := "Scott"

		err := testCache.Put(key,value)

		assert.NoError(t, err)

		assert.Equal(t, testCache.Read(key), value)
	})
}
