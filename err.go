package main

func Panic(e error) {
	if e != nil {
		panic(e)
	}
}
