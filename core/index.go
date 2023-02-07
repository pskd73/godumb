package core

import "os"

type IndexMap map[interface{}][]RecordAddress
type Order int

const (
	ASC  Order = 1
	DESC Order = 2
)

func GenerateIndexMap(readFile *os.File, key string) IndexMap {
	meta := GetDbMeta(readFile)
	indexMap := make(IndexMap)

	for i := int64(0); i < meta.Count; i++ {
		record := GetRecord(readFile, RecordAddress(i))
		UpdateIndexMap(indexMap, record[key], RecordAddress(i))
	}

	return indexMap
}

func UpdateIndexMap(indexMap IndexMap, key interface{}, idx RecordAddress) {
	if indexMap[key] == nil {
		indexMap[key] = []RecordAddress{}
	}
	indexMap[key] = append(indexMap[key], idx)
}

type Index struct {
	indexMap  IndexMap
	order     Order
	orderList []interface{}
}
