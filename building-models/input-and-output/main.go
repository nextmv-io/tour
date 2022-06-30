package main

import (
	"github.com/nextmv-io/sdk/run"
	"github.com/nextmv-io/sdk/store"
)

func main() {
	run.Run(
		func(input any, opt store.Options) (store.Solver, error) {
			return store.New().Satisfier(opt), nil
		},
	)
}
