package file

type VirtualBlockedFile struct {
	File BlockedFile
}

func (self VirtualBlockedFile) Get(addr int64) Block {
	block := self.File.Get(addr)
	return block
}

func (self VirtualBlockedFile) Push(data string) Block {
	block := self.File.Push(data)
	return block
}

func (self VirtualBlockedFile) All() []Block {
	blocks := self.File.All()
	return blocks
}

func (self VirtualBlockedFile) Set(addr int64, data string) {
	self.File.Set(addr, data)
}

func (self VirtualBlockedFile) Delete(addr int64) {
	self.Delete(addr)
}

func (self VirtualBlockedFile) Stat() FileStat {
	stat := self.File.Stat()
	return stat
}
