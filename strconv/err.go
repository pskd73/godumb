package strconv

func Panic(e error) {
	if e != nil {
		panic(e)
	}
}
