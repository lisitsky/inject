package main

import (
	"fmt"

	"github.com/lisitsky/inject"

	_ "github.com/lisitsky/inject/examples/simple/dependency"
)

var (
	str fmt.Stringer
)

func main() {
	inject.Construct(&str)
	fmt.Println("My Stringer is:", str)
}
