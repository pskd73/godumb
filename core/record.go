package core

type Record map[string]interface{}

func BlockToRecord(block string) Record {
	decoded := FromBase64(Unpad(block))
	return StringToJson(decoded)
}

func BlocksToRecords(blocks []string) []Record {
	records := []Record{}
	for _, block := range blocks {
		records = append(records, BlockToRecord(block))
	}
	return records
}

func RecordToBlock(record Record, blockSize int) string {
	jsonText := JsonToString(record)
	str, err := Pad(ToBase64(jsonText), blockSize)
	Panic(err)
	return str
}
