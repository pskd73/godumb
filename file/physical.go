package file

import (
	"bufio"
	"io/fs"
	"io/ioutil"
	"os"
)

type PhysicalFile struct {
	Path       string
	readFile   *os.File
	appendFile *os.File
}

func (self *PhysicalFile) Init() {
	readFile, err := os.Open(self.Path)
	Panic(err)
	self.readFile = readFile

	appendFile, err := os.OpenFile(self.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	Panic(err)
	self.appendFile = appendFile
}

func (self *PhysicalFile) Read(offset int64, limit int) []byte {
	self.readFile.Seek(offset, 0)

	r4 := bufio.NewReader(self.readFile)
	content, err := r4.Peek(limit)
	Panic(err)

	return content
}

func (self *PhysicalFile) Scan() []byte {
	self.readFile.Seek(0, 0)
	content, err := ioutil.ReadAll(self.readFile)
	Panic(err)
	return content
}

func (self *PhysicalFile) Append(data []byte) {
	_, err := self.appendFile.Write(data)
	Panic(err)
}

func (self *PhysicalFile) Stat() fs.FileInfo {
	info, err := self.readFile.Stat()
	Panic(err)
	return info
}
