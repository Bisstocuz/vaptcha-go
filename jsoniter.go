package vaptcha

import jsoniter "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func mustMarshalToJson(v interface{}) (result []byte) {
	result, _ = json.Marshal(v)
	return
}
