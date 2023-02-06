package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func createDummyDb(appendFile *os.File) {
	var record = map[string]string{
		"name": "Parmod",
	}

	for i := 0; i < 100000; i++ {
		fmt.Printf("%d\n", i)
		record["idx"] = strconv.Itoa(i)
		InsertRecord(appendFile, record)
	}
}

func main() {
	fileBuffers := GetFileBuffers("data.db")
	fmt.Println("Initiated")

	meta := GetDbMeta(fileBuffers.read)

	index := GenerateIndex(fileBuffers.read, "idx")
	fmt.Println("Generated index")

	t1 := time.Now()
	GetRecord(fileBuffers.read, index["99999"][0])
	t2 := time.Now()
	fmt.Println("With index: ", t2.Sub(t1))

	t1 = time.Now()
	for i := int64(0); i < meta.count; i++ {
		record := GetRecord(fileBuffers.read, i)
		if record["idx"] == "99999" {
			t2 = time.Now()
			fmt.Println("Without index: ", t2.Sub(t1))
		}
	}
}
