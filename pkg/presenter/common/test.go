package common

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
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
	require.Nil(t, err)
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	require.Nil(t, err)

	err = json.Unmarshal(byteValue, &result)
	require.Nil(t, err)
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
	require.Nil(t, err)
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	require.Nil(t, err)

	value := gjson.Get(string(byteValue), path)
	jsonResponses := make([]unmarshalString, 0)
	err = json.Unmarshal([]byte(value.String()), &jsonResponses)
	require.Nil(t, err)

	responses := make([]proto.Message, len(jsonResponses))
	for index, jsonResponse := range jsonResponses {
		var message proto.Message
		message = reflect.New(reflect.ValueOf(resultPtr).Elem().Type()).Interface().(proto.Message)
		err = protojson.Unmarshal(jsonResponse, message)
		require.Nil(t, err)
		responses[index] = message
	}

	return responses
}
