package main

import (
	"bufio"
	"os"
)

func GetReadFile(path string) *os.File {
	f, err := os.Open(path)
	Panic(err)
	return f
}

func GetAppendFile(path string) *os.File {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	Panic(err)
	return f
}

func ReadFile(file *os.File, offset int64, limit int) []byte {
	file.Seek(offset, 0)

	r4 := bufio.NewReader(file)
	content, err := r4.Peek(limit)
	Panic(err)

	return content
}

func Append(file *os.File, content []byte) {
	_, err := file.Write(content)
	Panic(err)
}
