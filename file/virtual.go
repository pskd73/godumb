package file

type VirtualBlockedFile struct {
	File         BlockedFile
	cachedBlocks map[int64]Block
}

func (self *VirtualBlockedFile) Init() {
	self.cachedBlocks = map[int64]Block{}
	self.pull()
}

func (self *VirtualBlockedFile) Get(addr int64) Block {
	if block, exists := self.cachedBlocks[addr]; exists {
		return block
	}
	block := self.File.Get(addr)
	return block
}

func (self *VirtualBlockedFile) Push(data string) Block {
	block := self.File.Push(data)
	self.cachedBlocks[block.Addr] = block
	return block
}

func (self *VirtualBlockedFile) All() []Block {
	blocks := []Block{}
	for _, block := range self.cachedBlocks {
		blocks = append(blocks, block)
	}
	return blocks
}

func (self *VirtualBlockedFile) Set(addr int64, data string) {
	self.File.Set(addr, data)
}

func (self *VirtualBlockedFile) Delete(addr int64) {
	self.Delete(addr)
}

func (self *VirtualBlockedFile) Stat() FileStat {
	stat := self.File.Stat()
	return stat
}

func (self *VirtualBlockedFile) pull() {
	blocks := self.File.All()
	for _, block := range blocks {
		self.cachedBlocks[block.Addr] = block
	}
}
