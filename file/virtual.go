package file

type VirtualFile struct {
	File          *PhysicalFile
	CachedContent string
}

func (self *VirtualFile) Init() {
	self.File.Init()
	self.CachedContent = string(self.File.Scan())
}

func (self *VirtualFile) Read(offset int64, limit int64) string {
	return self.CachedContent[offset : offset+limit]
}

func (self *VirtualFile) Scan() string {
	return self.CachedContent
}

func (self *VirtualFile) Append(data string) {
	self.CachedContent = self.CachedContent + data
	self.File.Append([]byte(data))
}
