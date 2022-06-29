package main

import (
	"fmt"

	"github.com/nextmv-io/sdk"
	"github.com/nextmv-io/sdk/store"
)

func main() {
	s := store.New()
	version := store.NewVar(s, sdk.VERSION)
	fmt.Println("Hello Hop", version.Get(s))
}
