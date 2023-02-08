package file

import (
	"fmt"
	"strings"
)

type FlexBlockedFile struct {
	File    *PhysicalFile
	MaxSize int
}

func (self FlexBlockedFile) Init() {
	self.File.Init()
}

func (self FlexBlockedFile) Get(addr int64) Block {
	probableString := string(self.File.Read(addr, self.MaxSize+1))
	parts := strings.Split(probableString, "\n")
	fmt.Println("flex_blocked", parts)
	return Block{Data: parts[0], Addr: addr}
}

func (self FlexBlockedFile) Push(data string) Block {
	nextAddr := self.Stat().Size + 1
	self.File.Append([]byte(data + "\n"))
	return Block{Data: data, Addr: nextAddr}
}

func (self FlexBlockedFile) All() []Block {
	lines := strings.Split(string(self.File.Scan()), "\n")
	blocks := []Block{}
	nextAddr := 0
	for _, line := range lines {
		blocks = append(blocks, Block{Data: line, Addr: int64(nextAddr)})
		nextAddr += len(line)
	}
	return blocks
}

func (self FlexBlockedFile) Set(idx int64, data string) {

}

func (self FlexBlockedFile) Delete(idx int64) {

}

func (self FlexBlockedFile) Stat() FileStat {
	stat := self.File.Stat()

	return FileStat{
		Count: stat.Size() / int64((self.MaxSize + 1)),
		Size:  stat.Size(),
	}
}
