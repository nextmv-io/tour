package main

import (
	"fmt"

	"github.com/nextmv-io/sdk/store"
)

func main() {
	s1 := store.New()
	x := store.NewVar(s1, 42)
	y := store.NewVar(s1, "foo")

	s2 := s1.Apply(y.Set("bar"))
	pi := store.NewVar(s2, 3.14)

	fmt.Println("s1:", x.Get(s1), y.Get(s1))
	fmt.Println("s2:", x.Get(s2), y.Get(s2), pi.Get(s2))
}
