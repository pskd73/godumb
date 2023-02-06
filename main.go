package main

import (
	"fmt"
	"time"
)

func main() {
	// appendFile := GetAppendFile("data")
	readFile := GetReadFile("data.db")

	// var record = map[string]string{
	// 	"name": "Parmod",
	// }

	// for i := 0; i < 10000000; i++ {
	// 	fmt.Printf("%d\n", i)
	// 	record["idx"] = strconv.Itoa(i)
	// 	InsertRecord(appendFile, record)
	// }

	// for i := int64(0); i < 10000; i++ {
	// 	record := GetRecord(readFile, i)
	// 	fmt.Printf("Compare %d %s\n", i, record["idx"].(string))
	// 	if record["idx"] != strconv.Itoa(int(i)) {
	// 		panic(i)
	// 	}
	// }

	t1 := time.Now()
	fmt.Println(GetRecord(readFile, 3823293))
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
}
