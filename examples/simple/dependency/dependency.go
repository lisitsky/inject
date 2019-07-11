package dependency

import (
	"fmt"

	"github.com/lisitsky/inject"
)

func init() {
	inject.Provide(NewStringer)
}

type stringer struct{}

func (s stringer) String() string {
	return "Hello, World"
}

func NewStringer() fmt.Stringer {
	return stringer{}
}
