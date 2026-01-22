package main

type DataSource interface {
	WriteData(data string)
	ReadData() string
}
