package file

func Panic(e error) {
	if e != nil {
		panic(e)
	}
}
