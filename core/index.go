package core

type IndexMap map[interface{}][]int64
type Order int

const (
	ASC  Order = 1
	DESC Order = 2
)

func GenerateIndexMap(records []Record, key string) IndexMap {
	indexMap := make(IndexMap)

	for i, record := range records {
		UpdateIndexMap(indexMap, record[key], int64(i))
	}

	return indexMap
}

func UpdateIndexMap(indexMap IndexMap, key interface{}, idx int64) {
	if indexMap[key] == nil {
		indexMap[key] = []int64{}
	}
	indexMap[key] = append(indexMap[key], idx)
}

type Index struct {
	indexMap  IndexMap
	order     Order
	orderList []interface{}
}
