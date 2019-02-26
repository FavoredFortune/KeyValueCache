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

func TestPut(t *testing.T) {

	t.Run("it can put and read", func(t *testing.T) {

		testCache := &SimpleKeyValueCache{map[string]string{}}
		require.NotNil(t, testCache)

		key := "testKey"
		value := "testValue"
		err := testCache.Put(key,value)

		assert.NoError(t,err)
		b, _ := testCache.Read(key)
		assert.Equal(t, b, value)

	})

	t.Run("second put test", func(t *testing.T) {
		testCache := &SimpleKeyValueCache{map[string]string{}}
		require.NotNil(t, testCache)
		key2 := "123"
		value2 := "Sooz"

		err2 := testCache.Put(key2, value2)
		assert.NoError(t, err2)

		a,_ := testCache.Read(key2)
		assert.Equal(t, a, value2)

	})
}


func TestRead(t *testing.T){
	t.Run("it can read", func(t *testing.T) {
		testCache := &SimpleKeyValueCache{map[string]string{}}
		require.NotNil(t, testCache)

		key := "name"
		value := "Scott"

		err := testCache.Put(key,value)

		assert.NoError(t, err)

		f, _ := testCache.Read(key)

		assert.Equal(t, f, value)
	})

	 t.Run("second read test", func(t *testing.T) {
		testCache := &SimpleKeyValueCache{map[string]string{}}
		require.NotNil(t, testCache)

		key := "name"
		value := "Scott"

		err := testCache.Put(key, value)
		assert.NoError(t, err)

		f, _ := testCache.Read(key)

		assert.Equal(t, f, value)

		_, err2 := testCache.Read("")

		//updated tests to reflect new Read method signature and used Objects are Equal values due to the indirect reference to the error message in the assertion
		assert.ObjectsAreEqualValues(err2, "read failed: key ' ' invalid")
	})
}
