package main

import (
	"encoding/json"
	"os"

	"github.com/nextmv-io/sdk/store"
)

func main() {
	enc := json.NewEncoder(os.Stdout)

	s := store.New()
	if err := enc.Encode(s); err != nil {
		panic(err)
	}

	x := store.NewVar(s, 42)
	y := store.NewVar(s, "foo")
	pi := store.NewVar(s, 3.14)
	if err := enc.Encode(s); err != nil {
		panic(err)
	}

	s = s.Format(func(s store.Store) any {
		return map[string]any{
			"x":  x.Get(s),
			"y":  y.Get(s),
			"pi": pi.Get(s),
		}
	})
	if err := enc.Encode(s); err != nil {
		panic(err)
	}

	if err := enc.Encode(s.Apply(y.Set("bar"))); err != nil {
		panic(err)
	}
}
