package main

import "fmt"

type FileDataSource struct {
	filename string
}

func NewFileDataSource(filename string) DataSource {
	return FileDataSource{filename: filename}
}

func (FileDataSource) WriteData(data string) {
	fmt.Println("Writing basic file: " + data)
}

func (FileDataSource) ReadData() string {
	return "basic file"
}
