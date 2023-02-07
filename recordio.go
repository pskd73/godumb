package main

import (
	"bytes"
	"os"
)

const RECORD_SIZE = 512

type DbMeta struct {
	count int64
	size  int64
}

func InsertRecord(appendFile *os.File, data interface{}) int64 {
	var content bytes.Buffer

	padded, err := Pad(ToBase64(JsonToString(data)), RECORD_SIZE)
	Panic(err)

	content.WriteString(padded)
	content.WriteString("\n")

	Append(appendFile, content.Bytes())

	return GetDbMeta(appendFile).count
}

func GetRecord(readFile *os.File, idx int64) map[string]interface{} {
	bytes := ReadFile(readFile, idx*(RECORD_SIZE+1), RECORD_SIZE)
	decoded := FromBase64(Unpad(string(bytes)))
	return StringToJson(decoded)
}

func GetDbMeta(readFile *os.File) DbMeta {
	stat, err := readFile.Stat()
	Panic(err)

	return DbMeta{
		count: stat.Size() / (RECORD_SIZE + 1),
		size:  stat.Size(),
	}
}
