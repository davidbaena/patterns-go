package main

type ConcreteFactoryNike struct{}

func (ConcreteFactoryNike) CreateShoe() IShoe {
	return NikeShoe{
		size:  39,
		color: "BLACK",
	}
}

func (ConcreteFactoryNike) CreateShirt() IShirt {
	return NikeShirt{
		size:      36,
		hasSlaves: true,
	}
}

func (ConcreteFactoryNike) CreateHat() IHat {
	return NikeHat{
		size:  60,
		color: "RED",
	}
}

func (ConcreteFactoryNike) CreateJacket() IJacket {
	return NikeJacket{
		size:          42,
		color:         "BLUE",
		isWatterProof: true,
	}
}

type NikeShoe struct {
	size  int
	color string
}

func (ns NikeShoe) Size() int {
	return ns.size
}

func (ns NikeShoe) Color() string {
	return ns.color
}

type NikeShirt struct {
	size      int
	hasSlaves bool
}

func (nsh NikeShirt) Size() int {
	return nsh.size
}

func (nsh NikeShirt) HasSleeves() bool {
	return nsh.hasSlaves
}

type NikeHat struct {
	size  int
	color string
}

func (nh NikeHat) Size() int {
	return nh.size
}

func (nh NikeHat) Color() string {
	return nh.color
}

type NikeJacket struct {
	size          int
	color         string
	isWatterProof bool
}

func (nj NikeJacket) Size() int {
	return nj.size
}

func (nj NikeJacket) Color() string {
	return nj.color
}

func (nj NikeJacket) IsWaterProof() bool {
	return nj.isWatterProof
}
