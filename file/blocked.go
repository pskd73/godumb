package file

import "strings"

type BlockedFileStat struct {
	Count int64
	Size  int64
}

type BlockedFile struct {
	File      *PhysicalFile
	BlockSize int
}

func (self *BlockedFile) Init() {
	self.File.Init()
}

func (self *BlockedFile) Get(idx int64) string {
	return string(self.File.Read(idx*int64(self.BlockSize+1), self.BlockSize))
}

func (self *BlockedFile) Push(data string) int64 {
	self.File.Append([]byte(data + "\n"))
	return self.Stat().Count - 1
}

func (self *BlockedFile) All() []string {
	return strings.Split(string(self.File.Scan()), "\n")
}

func (self *BlockedFile) Set(idx int64, data string) {

}

func (self *BlockedFile) Delete(idx int64) {

}

func (self *BlockedFile) Stat() BlockedFileStat {
	stat := self.File.Stat()

	return BlockedFileStat{
		Count: stat.Size() / int64((self.BlockSize + 1)),
		Size:  stat.Size(),
	}
}
