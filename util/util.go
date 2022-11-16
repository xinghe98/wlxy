package util

import (
	"encoding/json"
)

type Map[Key string, VALUE any] map[Key]VALUE

// json 转 map
func JsonToMap[T any](jsonStr string) Map[string, T] {
	var dat Map[string, T]
	if err := json.Unmarshal([]byte(jsonStr), &dat); err == nil {
		return dat
	}
	return nil
}

// map 转 json
func MapToJson[T string | int](m Map[string, T]) []byte {
	jsonStr, _ := json.Marshal(m)
	return jsonStr
}
