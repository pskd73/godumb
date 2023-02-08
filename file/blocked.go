package file

type FileStat struct {
	Count int64
	Size  int64
}

type Block struct {
	Data string
	Addr int64
}

type BlockedFile interface {
	Get(addr int64) Block
	Push(data string) Block
	All() []Block
	Set(addr int64, data string)
	Delete(addr int64)
	Stat() FileStat
}
