package main

import (
	"fmt"

	"godumb/core"
)

func insertRandomRecords(collection *core.Collection) {
	var record = map[string]interface{}{
		"name": "Parmod",
	}

	for i := 0; i < 100000; i++ {
		fmt.Printf("%d\n", i)
		collection.Insert(record)
	}
}

func main() {
	// collection := core.Collection{}
	// collection.Init("yelp_tip")
	// data, _ := ioutil.ReadFile("yelp_academic_dataset_tip.json")
	// lines := strings.Split(string(data), "\n")

	// for i, line := range lines {
	// 	fmt.Println(i)
	// 	var lineJson map[string]interface{}
	// 	json.Unmarshal([]byte(line), &lineJson)
	// 	if lineJson != nil {
	// 		collection.Insert(lineJson)
	// 	}
	// }

	// prompt.Run()

	coll := core.Collection{}
	coll.Init("ytl")
	coll.AddIndex("user_id")
	fmt.Println("Index created")
	fmt.Print(coll.GetByKey("date", "2013-02-28 02:05:54"))
	// t1 := time.Now()
	// for _, s := range coll.VirtualFile.CachedContent {
	// 	fmt.Print(s)
	// }
	// t2 := time.Now()

	// fmt.Println(t2.Sub(t1))
}
