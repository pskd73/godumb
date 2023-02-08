package core

type IndexMap map[interface{}][]int64
type Order int

const (
	ASC  Order = 1
	DESC Order = 2
)

func GenerateIndexMap(records []Record, key string) IndexMap {
	indexMap := make(IndexMap)

	for _, record := range records {
		UpdateIndexMap(indexMap, record.Data[key], record.Addr)
	}

	return indexMap
}

func UpdateIndexMap(indexMap IndexMap, key interface{}, addr int64) {
	if indexMap[key] == nil {
		indexMap[key] = []int64{}
	}
	indexMap[key] = append(indexMap[key], addr)
}

type Index struct {
	indexMap  IndexMap
	order     Order
	orderList []interface{}
}
