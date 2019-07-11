package pkgA

import "github.com/lisitsky/inject"

type Cat struct {
	Name string
}

func init() {
	inject.Provide(NewCat)
}

func NewCat(name string) *Cat {
	return &Cat{Name: name}
}
