package main

import (
	"fmt"
	"os"

	"github.com/google/uuid"
)

type Collection struct {
	name       string
	indexMaps  map[string]map[interface{}][]int64
	readFile   *os.File
	appendFile *os.File
}

func (self *Collection) Init(name string) {
	self.name = name
	self.readFile = GetReadFile(fmt.Sprintf("%s.db", name))
	self.appendFile = GetAppendFile(fmt.Sprintf("%s.db", name))
	self.indexMaps = make(map[string]map[interface{}][]int64)
	self.indexMaps["_id"] = GenerateIndexMap(self.readFile, "_id")
	fmt.Printf("Collection: %s initiated\n", name)
}

func (self *Collection) GetRecordByIdx(idx int64) map[string]interface{} {
	return GetRecord(self.readFile, idx)
}

func (self *Collection) GetRecordById(id string) map[string]interface{} {
	return self.GetRecordByIdx(self.indexMaps["_id"][id][0])
}

func (self *Collection) Insert(record map[string]interface{}) {
	record["_id"] = uuid.New().String()
	idx := InsertRecord(self.appendFile, record)
	UpdateIndexMap(self.indexMaps["_id"], record["_id"], idx)
}

func (self *Collection) GetMeta() DbMeta {
	return GetDbMeta(self.readFile)
}

func (self *Collection) GetAllRecords() []map[string]interface{} {
	meta := self.GetMeta()
	var records []map[string]interface{}
	for i := int64(0); i < meta.count; i++ {
		records = append(records, self.GetRecordByIdx(i))
	}
	return records
}
