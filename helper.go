package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

func ToBase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func FromBase64(str string) string {
	decoded, err := base64.StdEncoding.DecodeString(str)
	Panic(err)
	return string(decoded)
}

func JsonToString(jsonObject interface{}) string {
	jsonBytes, err := json.Marshal(jsonObject)
	Panic(err)
	return string(jsonBytes)
}

func StringToJson(str string) map[string]interface{} {
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(str), &jsonMap)
	return jsonMap
}

func Pad(str string, maxSize int) (string, error) {
	if len(str) > maxSize {
		return "", errors.New("str cannot be more than maxSize")
	}

	var buf strings.Builder

	buf.WriteString(str)
	buf.WriteString(strings.Repeat("#", maxSize-len(str)))

	return buf.String(), nil
}

func Unpad(str string) string {
	return strings.ReplaceAll(str, "#", "")
}
