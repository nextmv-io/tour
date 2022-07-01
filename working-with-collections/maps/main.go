package main

import (
	"fmt"

	"github.com/nextmv-io/sdk/store"
)

func main() {
	s1 := store.New()
	x := store.NewMap[string, float64](s1)

	s2 := s1.Apply(
		x.Set("pi", 3.14),
		x.Set("e", 2.72),
	)

	s3 := s2.Apply(x.Delete("e"))

	fmt.Println(x.Map(s1))
	fmt.Println(x.Map(s2))
	fmt.Println(x.Map(s3))
}
