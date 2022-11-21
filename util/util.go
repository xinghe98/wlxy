package util

import (
	jsoniter "github.com/json-iterator/go"
	"io"
	"strconv"
	"time"
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

// ResolveTime 将秒数转换为时分秒
func ResolveTime(seconds int) (hour, minute, second int) {
	var day = seconds / (24 * 3600)
	hour = (seconds - day*3600*24) / 3600
	minute = (seconds - day*24*3600 - hour*3600) / 60
	second = seconds - day*24*3600 - hour*3600 - minute*60
	return hour, minute, second
}

// GenerateTime 时间格式化,生成一个start_time
func GenerateTime(Times int) (timestr string) {
	now := time.Now()
	s, _ := time.ParseDuration("-" + strconv.Itoa(Times) + "s")
	t := now.Add(s)
	timestr = t.Format("2006-01-02 15:04:05")
	return timestr
}
