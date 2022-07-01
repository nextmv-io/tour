package main

import (
	"encoding/json"
	"os"

	"github.com/nextmv-io/sdk/model"
	"github.com/nextmv-io/sdk/store"
)

func main() {
	s1 := store.New()
	d1 := store.NewDomain(s1)
	d2 := store.NewDomain(s1, model.NewRange(1, 10))
	d3 := store.NewDomain(s1, model.NewRange(-5, 5), model.NewRange(15, 25))

	s2 := s1.Apply(
		d1.Add(42, 43, 44),
		d2.Remove(2, 4, 6, 8, 10),
		d3.AtLeast(10),
	)

	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(s2); err != nil {
		panic(err)
	}

	s3 := store.New()
	d := store.Repeat(s3, 5, model.NewDomain(model.NewRange(1, 10)))

	s4 := s3.Apply(
		d.Add(0, 42),
		d.Assign(2, 5),
		d.Remove(4, 2, 3, 4),
	)

	if err := enc.Encode(s4); err != nil {
		panic(err)
	}
}
