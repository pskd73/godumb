package core

import (
	"fmt"
	"os"

	"github.com/google/uuid"
)

type Collection struct {
	name       string
	indexMaps  map[string]IndexMap
	readFile   *os.File
	appendFile *os.File
}

func (self *Collection) Init(name string) {
	self.name = name
	self.readFile = GetReadFile(fmt.Sprintf("%s.db", name))
	self.appendFile = GetAppendFile(fmt.Sprintf("%s.db", name))
	self.indexMaps = make(map[string]IndexMap)
	self.indexMaps["_id"] = GenerateIndexMap(self.readFile, "_id")
	fmt.Printf("Collection: %s initiated\n", name)
}

func (self *Collection) GetRecordByIdx(idx RecordAddress) Record {
	return GetRecord(self.readFile, idx)
}

func (self *Collection) GetRecordById(id string) Record {
	return self.GetRecordByIdx(self.indexMaps["_id"][id][0])
}

func (self *Collection) Insert(record Record) {
	record["_id"] = uuid.New().String()
	idx := InsertRecord(self.appendFile, record)
	UpdateIndexMap(self.indexMaps["_id"], record["_id"], idx)
}

func (self *Collection) GetMeta() DbMeta {
	return GetDbMeta(self.readFile)
}

func (self *Collection) GetAllRecords() []Record {
	meta := self.GetMeta()
	var records []Record
	for i := int64(0); i < meta.Count; i++ {
		records = append(records, self.GetRecordByIdx(RecordAddress(i)))
	}
	return records
}

func (self *Collection) AddIndex(field string) {
	self.indexMaps[field] = GenerateIndexMap(self.readFile, field)
}

func (self *Collection) GetByKey(key string, value string) []Record {
	var addresses []RecordAddress
	if indexMap, exists := self.indexMaps[key]; exists {
		if addrs, rowExists := indexMap[value]; rowExists {
			addresses = addrs
		}
	}
	records := []Record{}
	if addresses != nil {
		for _, addr := range addresses {
			records = append(records, self.GetRecordByIdx(addr))
		}
		return records
	}
	for i := int64(0); i < self.GetMeta().Count; i++ {
		record := self.GetRecordByIdx(RecordAddress(i))
		if record[key] == value {
			records = append(records, record)
		}
	}
	return records
}
