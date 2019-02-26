package kvcache

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSimpleKeyValueCache(t *testing.T) {
	t.Run("new cache created", func(t *testing.T) {
		cache := NewSimpleKVCache()
		assert.NotNil(t, cache)
	})

}

func TestPut(t *testing.T) {


	t.Run("it can put and read", func(t *testing.T) {

		cache := NewSimpleKVCache()
		require.NotNil(t, cache)

		key := "testKey"
		value := "testValue"
		err := cache.Put(key,value)

		assert.NoError(t,err)
		assert.ObjectsAreEqualValues(cache.Read(key),value)

	})

}