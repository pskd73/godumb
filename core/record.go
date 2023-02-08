package core

import (
	"godumb/strconv"

	"godumb/file"
)

type Record struct {
	Data map[string]interface{}
	Addr int64
}

func BlockToRecord(block file.Block) Record {
	decoded := strconv.FromBase64(block.Data)
	return Record{Data: strconv.StringToJson(decoded), Addr: block.Addr}
}

func BlocksToRecords(blocks []file.Block) []Record {
	records := []Record{}
	for _, block := range blocks {
		records = append(records, BlockToRecord(block))
	}
	return records
}

func DataToString(data map[string]interface{}) string {
	jsonText := strconv.JsonToString(data)
	return strconv.ToBase64(jsonText)
}
