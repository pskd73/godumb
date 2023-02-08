package prompt

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"godumb/core"
)

type State struct {
	Collection *core.Collection
}

type MenuItem struct {
	Name string
	Func func(*State)
}

func PrintMenu(menu []MenuItem) {
	for k, v := range menu {
		fmt.Printf("%d. %s\n", k+1, v.Name)
	}
}

func LoadCollection(state *State) {
	var name string
	fmt.Printf("Enter collection name: ")
	fmt.Scan(&name)
	state.Collection = &core.Collection{}
	state.Collection.Init(name)
}

func Exit(state *State) {
	os.Exit(0)
}

func GetById(state *State) {
	var id string
	fmt.Printf("Enter _id: ")
	fmt.Scan(&id)
	fmt.Println(state.Collection.GetRecordById(id))
}

func GetByAddr(state *State) {
	var addr string
	fmt.Printf("Enter idx: ")
	fmt.Scan(&addr)

	idxInt, err := strconv.Atoi(addr)
	if err != nil {
		return
	}
	fmt.Println(state.Collection.GetRecordByAddr(int64(idxInt)))
}

func AddIndex(state *State) {
	var field string
	fmt.Printf("Enter key: ")
	fmt.Scan(&field)
	state.Collection.AddIndex(field)
}

func GetByKey(state *State) {
	var key string
	var val string
	fmt.Printf("Enter key: ")
	fmt.Scan(&key)
	fmt.Printf("Enter value: ")
	fmt.Scan(&val)
	t1 := time.Now()
	fmt.Println(state.Collection.GetByKey(key, val))
	t2 := time.Now()
	fmt.Printf("Time taken: %s", t2.Sub(t1))
}

func Run() {
	menu := []MenuItem{
		{Name: "Load collection", Func: LoadCollection},
		{Name: "Add index", Func: AddIndex},
		{Name: "Get by idx", Func: GetByAddr},
		{Name: "Get by _id", Func: GetById},
		{Name: "Get by key", Func: GetByKey},
		{Name: "Exit", Func: Exit},
	}

	state := State{}
	// state.Collection = &core.Collection{}
	// state.Collection.Init("yelp_tip")

	for true {
		var input string
		fmt.Printf("\n\n")
		PrintMenu(menu)
		fmt.Printf("\nEnter the menu number: ")
		fmt.Scan(&input)

		if inputInt, err := strconv.Atoi(input); err != nil {
			continue
		} else {
			if inputInt < 1 || inputInt > len(menu) {
				continue
			}
			menu[inputInt-1].Func(&state)
		}
	}

	PrintMenu(menu)
}
