package main

import (
	"fmt"
	"time"
)

func insertRandomRecords(collection *Collection) {
	var record = map[string]interface{}{
		"name": "Parmod",
	}

	for i := 0; i < 100000; i++ {
		fmt.Printf("%d\n", i)
		collection.Insert(record)
	}
}

func main() {
	collection := Collection{}
	collection.Init("data")
	// insertRandomRecords(&collection)
	// fmt.Println(collection.GetMeta().count)
	// for _, record := range collection.GetAllRecords() {
	// 	fmt.Println(record)
	// }

	t1 := time.Now()
	collection.GetRecordById("dadcfbb7-7651-41f6-b5ec-fac18a673c00")
	t2 := time.Now()
	fmt.Println("With index: ", t2.Sub(t1))

	t1 = time.Now()
	for i := int64(0); i < collection.GetMeta().count; i++ {
		record := collection.GetRecordByIdx(i)
		if record["_id"] == "dadcfbb7-7651-41f6-b5ec-fac18a673c00" {
			t2 = time.Now()
			fmt.Println("Without index: ", t2.Sub(t1))
		}
	}
}
