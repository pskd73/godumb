package file

import (
	"strings"

	"godumb/strconv"
)

type FixedBlockedFile struct {
	File      PhysicalFile
	BlockSize int
}

func (self *FixedBlockedFile) Init() {
	self.File.Init()
}

func (self *FixedBlockedFile) Get(addr int64) Block {
	offset := addr * (int64(self.BlockSize) + 1)
	unpadded := strconv.Unpad(string(self.File.Read(offset, self.BlockSize)))
	return Block{
		Data: unpadded,
		Addr: addr,
	}
}

func (self *FixedBlockedFile) Push(data string) Block {
	nextAddr := self.Stat().Count * int64(self.BlockSize+1)
	padded, err := strconv.Pad(data, self.BlockSize)
	Panic(err)
	self.File.Append([]byte(padded + "\n"))
	return Block{Data: data, Addr: nextAddr}
}

func (self *FixedBlockedFile) All() []Block {
	lines := strings.Split(string(self.File.Scan()), "\n")
	blocks := []Block{}
	for i, line := range lines {
		unpadded := strconv.Unpad(line)
		blocks = append(blocks, Block{Data: unpadded, Addr: int64(i)})
	}
	return blocks
}

func (self *FixedBlockedFile) Set(idx int64, data string) {

}

func (self *FixedBlockedFile) Delete(idx int64) {

}

func (self *FixedBlockedFile) Stat() FileStat {
	stat := self.File.Stat()

	return FileStat{
		Count: stat.Size() / int64((self.BlockSize + 1)),
		Size:  stat.Size(),
	}
}
