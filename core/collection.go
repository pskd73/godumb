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

	physicalFile := file.PhysicalFile{Path: fmt.Sprintf("%s.db", name)}
	blockedFile := file.FixedBlockedFile{File: physicalFile, BlockSize: BLOCK_SIZE}
	blockedFile.Init()

	virtualFile := file.VirtualBlockedFile{File: &blockedFile}
	virtualFile.Init()

	self.File = &virtualFile

	self.IndexMaps = make(map[string]IndexMap)
	self.IndexMaps["_id"] = GenerateIndexMap(BlocksToRecords(self.File.All()), "_id")
	fmt.Printf("Collection: %s initiated\n", name)
}

func (self *Collection) GetRecordByAddr(addr int64) Record {
	return BlockToRecord(self.File.Get(addr))
}

func (self *Collection) GetRecordById(id string) Record {
	return self.GetRecordByAddr(self.IndexMaps["_id"][id][0])
}

func (self *Collection) Insert(data map[string]interface{}) {
	data["_id"] = uuid.New().String()
	block := self.File.Push(DataToString(data))
	UpdateIndexMap(self.IndexMaps["_id"], data["_id"], block.Addr)
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
			records = append(records, self.GetRecordByAddr(addr))
		}
		return records
	}
	for _, record := range self.GetAllRecords() {
		if record.Data[key] == value {
			records = append(records, record)
		}
	}
	return records
}
