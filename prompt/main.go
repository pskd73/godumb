package prompt

import (
	"fmt"
	"os"

	"godumb/core"
)

type State struct {
	Collection *core.Collection
}

type MenuItem struct {
	Name string
	Func func(*State)
}

func PrintMenu(menu map[string]MenuItem) {
	for k, v := range menu {
		fmt.Printf("%s. %s\n", k, v.Name)
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
	fmt.Println(state.Collection.GetByKey(key, val))
}

func Run() {
	menu := map[string]MenuItem{
		"1": {Name: "Load collection", Func: LoadCollection},
		"2": {Name: "Exit", Func: Exit},
		"3": {Name: "Get by _id", Func: GetById},
		"4": {Name: "Get by key", Func: GetByKey},
		"5": {Name: "Add index", Func: AddIndex},
	}

	state := State{}

	for true {
		var input string
		fmt.Printf("\n\n")
		PrintMenu(menu)
		fmt.Printf("\nEnter the menu number: ")
		fmt.Scan(&input)
		if menuItem, ok := menu[input]; ok {
			menuItem.Func(&state)
		} else {
			fmt.Printf("Please enter valid option!\n")
		}
	}

	PrintMenu(menu)
}
