package util

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// JsontoMap Json转map
func JsonToMap(Body io.ReadCloser) map[string]interface{} {
	var response map[string]interface{}
	body, _ := ioutil.ReadAll(Body)
	err := json.Unmarshal(body, &response)
	if err != nil {
		return nil
	}
	return response
}
