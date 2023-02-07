package main

import "os"

func GenerateIndexMap(readFile *os.File, key string) map[interface{}][]int64 {
	meta := GetDbMeta(readFile)
	indexMap := make(map[interface{}][]int64)

	for i := int64(0); i < meta.count; i++ {
		record := GetRecord(readFile, i)
		UpdateIndexMap(indexMap, record[key], i)
	}

	return indexMap
}

func UpdateIndexMap(indexMap map[interface{}][]int64, key interface{}, idx int64) {
	if indexMap[key] == nil {
		indexMap[key] = []int64{}
	}
	indexMap[key] = append(indexMap[key], idx)
}
