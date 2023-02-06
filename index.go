package main

import "os"

func GenerateIndex(readFile *os.File, field string) map[interface{}][]int64 {
	meta := GetDbMeta(readFile)
	index := make(map[interface{}][]int64)

	for i := int64(0); i < meta.count; i++ {
		record := GetRecord(readFile, i)
		if index[record[field]] == nil {
			index[record[field]] = []int64{}
		}
		index[record[field]] = append(index[record[field]], i)
	}

	return index
}
