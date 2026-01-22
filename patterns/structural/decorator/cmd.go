package main

func main() {
	source := NewFileDataSource("file.txt")
	source.WriteData("data")

	source = NewEncryptionDecorator(source)
	source.WriteData("data")

	source = NewCompressionDecorator(source)
	source.WriteData("data")

}
