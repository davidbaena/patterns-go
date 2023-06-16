package main

type EncryptionDecorator struct {
	wrappee DataSource
}

func NewEncryptionDecorator(wrappee DataSource) DataSource {
	return EncryptionDecorator{wrappee: wrappee}
}

func (ed EncryptionDecorator) WriteData(data string) {
	data = data + "-encoded"
	ed.wrappee.WriteData(data)
}

func (ed EncryptionDecorator) ReadData() string {
	return ed.wrappee.ReadData()
}
