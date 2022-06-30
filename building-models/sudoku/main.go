package main

import (
	"github.com/nextmv-io/sdk/model"
	"github.com/nextmv-io/sdk/run"
	"github.com/nextmv-io/sdk/store"
)

func main() {
	run.Run(sudoku)
}

func sudoku(input [9][9]int, opt store.Options) (store.Solver, error) {
	root := store.New()
	x := store.Repeat(root, 9*9, model.NewDomain(model.NewRange(1, 9)))

	for i, row := range input {
		for j, cell := range row {
			if cell >= 1 && cell <= 9 {
				root = assign(root, x, i, j, cell)
			}
		}
	}

	root = root.Generate(
		store.Scope(
			func(s store.Store) store.Generator {
				i, ok := x.Smallest(s)
				values := x.Domain(s, i).Slice()

				for i := 0; i < 9*9; i++ {
					if x.Domain(s, i).Empty() {
						ok = false
						break
					}
				}

				return store.If(
					func(_ store.Store) bool {
						return ok && len(values) > 0
					},
				).Then(
					func(_ store.Store) store.Store {
						next := values[0]
						values = values[1:]
						return assign(s, x, i/9, i%9, next)
					},
				).With(x.Singleton)
			},
		),
	).Format(
		func(s store.Store) any {
			grid := [9][9]model.Domain{}
			for i := 0; i < 9*9; i++ {
				grid[i/9][i%9] = x.Domain(s, i)
			}
			return grid
		},
	)

	return root.Satisfier(opt), nil
}

func assign(s store.Store, x store.Domains, row, col, value int) store.Store {
	ind := index(row, col)
	changes := []store.Change{x.Assign(ind, value)}

	for j := 0; j < 9; j++ {
		if j != col {
			changes = append(changes, x.Remove(index(row, j), value))
		}
		if j != row {
			changes = append(changes, x.Remove(index(j, col), value))
		}
	}

	i, j := 3*(row/3), 3*(col/3)
	for m := i; m < i+3; m++ {
		for n := j; n < j+3; n++ {
			if k := index(m, n); k != ind {
				changes = append(changes, x.Remove(k, value))
			}
		}
	}

	return s.Apply(changes...)
}

func index(row, col int) int {
	return (row * 9) + col
}
