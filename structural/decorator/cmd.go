package main

func main() {
	source := NewFileDataSource("file.txt")
	source.WriteData("data")

	source = NewEncryptionDecorator(source)
	source.WriteData("data_to_encrypt")

	source = NewCompressionDecorator(source)
	source.WriteData("data_to_compress")

}
