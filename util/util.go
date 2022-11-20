package util

import (
	jsoniter "github.com/json-iterator/go"
	"io"
)

// JsontoMap Json转map
func JsonToMap(Body io.Reader) map[string]interface{} {
	var response map[string]interface{}
	var jsonMarshal = jsoniter.ConfigCompatibleWithStandardLibrary
	err := jsonMarshal.NewDecoder(Body).Decode(&response)
	if err != nil {
		return nil
	}
	return response
}
