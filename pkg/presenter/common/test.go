package common

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

const DefaultDepFixtureFile = "dep_fixture_test.json"

func TestDepPresenters(t *testing.T, result interface{}, filepath string) {
	if filepath == "" {
		filepath = DefaultDepFixtureFile
	}

	jsonFile, err := os.Open(filepath)
	if !assert.Nil(t, err) {
		t.Fail()
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if !assert.Nil(t, err) {
		t.Fail()
	}

	err = json.Unmarshal(byteValue, &result)
	if !assert.Nil(t, err) {
		t.Fail()
	}
}

type unmarshalString []byte

func (v *unmarshalString) UnmarshalJSON(data []byte) error {
	*v = data
	return nil
}

func TestEmailResponses(t *testing.T, resultPtr proto.Message, filepath, path string) []proto.Message {
	if filepath == "" {
		filepath = DefaultDepFixtureFile
	}

	if path == "" {
		path = "@this"
	}

	jsonFile, err := os.Open(filepath)
	if !assert.Nil(t, err) {
		t.Fail()
		return nil
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if !assert.Nil(t, err) {
		t.Fail()
		return nil
	}

	value := gjson.Get(string(byteValue), path)
	jsonResponses := make([]unmarshalString, 0)
	err = json.Unmarshal([]byte(value.String()), &jsonResponses)
	if !assert.Nil(t, err) {
		t.Fail()
		return nil
	}

	responses := make([]proto.Message, len(jsonResponses))
	for index, jsonResponse := range jsonResponses {
		var message proto.Message
		message = reflect.New(reflect.ValueOf(resultPtr).Elem().Type()).Interface().(proto.Message)
		err = protojson.Unmarshal(jsonResponse, message)
		if !assert.Nil(t, err) {
			t.Fail()
			return nil
		}
		responses[index] = message
	}

	return responses
}
