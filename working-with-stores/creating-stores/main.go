package main

import (
	"fmt"

	"github.com/nextmv-io/sdk/store"
)

func main() {
	s := store.New()

	x := store.NewVar(s, 42)
	y := store.NewVar(s, []float64{3.14, 2.72})

	fmt.Println(
		x.Get(s)*10,
		y.Get(s)[0],
	)
}
