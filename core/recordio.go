package core

import (
	"bytes"
	"os"
)

const RECORD_SIZE = 1 * 1024

type Record map[string]interface{}
type RecordAddress int64

type DbMeta struct {
	Count int64
	Size  int64
}

func InsertRecord(appendFile *os.File, data interface{}) RecordAddress {
	var content bytes.Buffer

	padded, err := Pad(ToBase64(JsonToString(data)), RECORD_SIZE)
	Panic(err)

	content.WriteString(padded)
	content.WriteString("\n")

	Append(appendFile, content.Bytes())

	return RecordAddress(GetDbMeta(appendFile).Count)
}

func GetRecord(readFile *os.File, idx RecordAddress) Record {
	bytes := ReadFile(readFile, int64(idx*(RECORD_SIZE+1)), RECORD_SIZE)
	decoded := FromBase64(Unpad(string(bytes)))
	return StringToJson(decoded)
}

func GetDbMeta(readFile *os.File) DbMeta {
	stat, err := readFile.Stat()
	Panic(err)

	return DbMeta{
		Count: stat.Size() / (RECORD_SIZE + 1),
		Size:  stat.Size(),
	}
}
