package main

type CompressionDecorator struct {
	wrappee DataSource
}

func NewCompressionDecorator(wrappee DataSource) DataSource {
	return CompressionDecorator{wrappee: wrappee}
}

func (ed CompressionDecorator) WriteData(data string) {
	data = data + "-compress"
	ed.wrappee.WriteData(data)
}

func (ed CompressionDecorator) ReadData() string {
	return ed.wrappee.ReadData()
}
