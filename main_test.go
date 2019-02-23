package main

import "testing"

func main() {

}

type TestKeyValueCache interface {
	Put(key, value string) error
	Read(key string) string
	Update(key, value string) error
	Delete(key string) error
}

type TestSimpleKeyValueCache struct {
	data map[string]string
}

var testCache = &TestSimpleKeyValueCache{map[string]string{}}

var testInputHappy = `
PUT name Sooz
READ name
UPDATE name Betty
READ name
DELETE name
`
var testOutputHappy = `
Sooz
Betty`

func Test_simpleKeyValueCache_Put(t *testing.T) {
	type fields struct {
		data map[string]string
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &simpleKeyValueCache{
				data: tt.fields.data,
			}
			if err := c.Put(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("simpleKeyValueCache.Put() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
