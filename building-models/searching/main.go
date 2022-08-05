package main

import (
	"errors"

	"github.com/nextmv-io/sdk/model"
	"github.com/nextmv-io/sdk/run"
	"github.com/nextmv-io/sdk/store"
)

func main() {
	run.Run(handler)
}

func handler(n int, opt store.Options) (store.Solver, error) {
	if n < 1 {
		return nil, errors.New("input must be > 1")
	}

	root := store.New()
	unused := store.NewDomain(root, model.NewRange(1, n))
	permutation := store.NewSlice[int](root)

	root = root.
		Validate(unused.Empty).
		Format(
			func(s store.Store) any {
				return permutation.Slice(s)
			},
		).
		Generate(func(s store.Store) store.Generator {
			values := unused.Slice(s)
			return store.Lazy(
				func() bool { return len(values) > 0 },
				func() store.Store {
					next := values[0]
					values = values[1:]

					return s.Apply(
						unused.Remove(next),
						permutation.Append(next),
					)
				},
			)
		})

	return root.Satisfier(opt), nil
}
