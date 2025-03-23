package main

type ConcreteFactoryAdidas struct{}

func (ConcreteFactoryAdidas) CreateShoe() IShoe {
	return AdidasShoe{
		size:  30,
		color: "BLUE",
	}
}

func (ConcreteFactoryAdidas) CreateShirt() IShirt {
	return AdidasShirt{
		size:      34,
		hasSlaves: true,
	}
}

func (ConcreteFactoryAdidas) CreateHat() IHat {
	return AdidasHat{
		size:  58,
		color: "WHITE",
	}
}

type AdidasShoe struct {
	size  int
	color string
}

func (ns AdidasShoe) Size() int {
	return ns.size
}

func (ns AdidasShoe) Color() string {
	return ns.color
}

type AdidasShirt struct {
	size      int
	hasSlaves bool
}

func (nsh AdidasShirt) Size() int {
	return nsh.size
}

func (nsh AdidasShirt) HasSleeves() bool {
	return nsh.hasSlaves
}

type AdidasHat struct {
	size  int
	color string
}

func (ah AdidasHat) Size() int {
	return ah.size
}

func (ah AdidasHat) Color() string {
	return ah.color
}
