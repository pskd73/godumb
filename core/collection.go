package core

import (
	"fmt"

	"godumb/file"

	"github.com/google/uuid"
)

const BLOCK_SIZE = 1024 * 1

type Collection struct {
	name      string
	IndexMaps map[string]IndexMap
	File      file.BlockedFile
}

func (self *Collection) Init(name string) {
	self.name = name

	f := file.PhysicalFile{Path: fmt.Sprintf("%s.db", name)}
	file := file.BlockedFile{File: &f, BlockSize: BLOCK_SIZE}
	file.Init()
	self.File = file
	self.IndexMaps = make(map[string]IndexMap)
	self.IndexMaps["_id"] = GenerateIndexMap(BlocksToRecords(self.File.All()), "_id")
	fmt.Printf("Collection: %s initiated\n", name)
}

func (self *Collection) GetRecordByIdx(idx int64) Record {
	return BlockToRecord(self.File.Get(idx))
}

func (self *Collection) GetRecordById(id string) Record {
	return self.GetRecordByIdx(self.IndexMaps["_id"][id][0])
}

func (self *Collection) Insert(record Record) {
	record["_id"] = uuid.New().String()
	idx := self.File.Push(RecordToBlock(record, self.File.BlockSize))
	UpdateIndexMap(self.IndexMaps["_id"], record["_id"], idx)
}

func (self *Collection) GetAllRecords() []Record {
	return BlocksToRecords(self.File.All())
}

func (self *Collection) AddIndex(field string) {
	self.IndexMaps[field] = GenerateIndexMap(self.GetAllRecords(), field)
}

func (self *Collection) GetByKey(key string, value interface{}) []Record {
	var addresses []int64
	if indexMap, exists := self.IndexMaps[key]; exists {
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
	for _, record := range self.GetAllRecords() {
		if record[key] == value {
			records = append(records, record)
		}
	}
	return records
}
