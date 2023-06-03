package main

import "fmt"

type FactoryType string

const (
	NikeFactory   FactoryType = "Nike"
	AdidasFactory FactoryType = "Adidas"
)

type ISportsFactory interface {
	CreateShoe() IShoe
	CreateShirt() IShirt
}

func ProvideSportFactory(brand FactoryType) (ISportsFactory, error) {
	switch brand {
	case NikeFactory:
		return ConcreteFactoryNike{}, nil
	case AdidasFactory:
		return ConcreteFactoryAdidas{}, nil
	default:
		return nil, fmt.Errorf("ErrWrongBrand")
	}
}
