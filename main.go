package main

import (
	"fmt"

	"godumb/core"
	"godumb/prompt"
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
	// 	if i > 1000000000000000 {
	// 		break
	// 	}
	// }

	prompt.Run()

	// coll := core.Collection{}
	// coll.Init("ytl")
	// coll.AddIndex("user_id")
	// fmt.Println("Index created")
	// fmt.Println(coll.IndexMaps)
	// fmt.Println(coll.GetByKey("_id", "d5bbdef7-5c37-4dc4-aa33-556a160fd7f5"))
	// fmt.Print(coll.GetByKey("user_id", "HAApeWFR7aBy0OGT2Y4Qvg"))
	// fmt.Println(coll.GetRecordByAddr(1)) //740166
	// t1 := time.Now()
	// for _, s := range coll.VirtualFile.CachedContent {
	// 	fmt.Print(s)
	// }
	// t2 := time.Now()

	// fmt.Println(t2.Sub(t1))
}
