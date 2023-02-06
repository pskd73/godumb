package main

import (
	"bytes"
	"os"
)

const RECORD_SIZE = 512

func InsertRecord(appendFile *os.File, data interface{}) {
	var content bytes.Buffer

	padded, err := Pad(ToBase64(JsonToString(data)), RECORD_SIZE)
	Panic(err)

	content.WriteString(padded)
	content.WriteString("\n")

	Append(appendFile, content.Bytes())
}

func GetRecord(readFile *os.File, idx int64) map[string]interface{} {
	bytes := ReadFile(readFile, idx*(RECORD_SIZE+1), RECORD_SIZE)
	decoded := FromBase64(Unpad(string(bytes)))
	return StringToJson(decoded)
}
