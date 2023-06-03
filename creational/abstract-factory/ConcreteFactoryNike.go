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
