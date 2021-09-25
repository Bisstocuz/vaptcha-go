package vaptcha

import jsoniter "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func mustToJson(v interface{}) (result []byte) {
	result, _ = json.Marshal(v)
	return
}
