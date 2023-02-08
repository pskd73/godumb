package main

import (
	"fmt"

	"godumb/core"
	"godumb/util"
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

type Human interface {
	SetName(string)
	GetName() string
}

type Person struct {
	Name string
}

func (self *Person) SetName(name string) {
	self.Name = name
}

func (self *Person) GetName() string {
	return self.Name
}

type World struct {
	Human
}

func (self World) GetHumanName() string {
	return self.Human.GetName()
}

func main() {
	timer := util.Timer{}
	timer.Init()

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

	// prompt.Run()

	// person := Person{}
	// person.SetName("Pramod")

	// var world World = World{Human: &person}
	// world.Human = &person
	// fmt.Println(world.GetHumanName())

	coll := core.Collection{}
	coll.Init("yelp_tip")
	// coll.AddIndex("user_id")
	// fmt.Println("Index created")
	// fmt.Println(coll.IndexMaps)

	timer.Check(false)
	coll.GetByKey("business_id", "Z9Q1IgIPB64Um_BNWpWgCA")
	timer.Check(true)

	coll.AddIndex("business_id")

	timer.Check(false)
	coll.GetByKey("business_id", "Z9Q1IgIPB64Um_BNWpWgCA")
	timer.Check(true)
	// fmt.Print(coll.GetByKey("user_id", "HAApeWFR7aBy0OGT2Y4Qvg"))
	// fmt.Println(coll.GetRecordByAddr(1)) //740166
	// t1 := time.Now()
	// for _, s := range coll.VirtualFile.CachedContent {
	// 	fmt.Print(s)
	// }
	// t2 := time.Now()

	// fmt.Println(t2.Sub(t1))
}
