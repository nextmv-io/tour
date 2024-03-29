# ![ears](../img/ears.png) Taking the Tour

To take this tour, you can either `git clone git@github.com:nextmv-io/tour.git`
or create a new module and start from scratch. Either way will work, but the
following instructions assume you are starting from scratch. This is so you get
used to building models the right way. If you'd prefer to skip some steps and
clone the tour instead, you should be able to follow along and `nextmv sdk run`
any of the `main.go` files included in the source tree.

You have the following resources available to you:

- [Nextmv SDK source](https://github.com/nextmv-io/sdk)
- [Nextmv SDK docs](https://pkg.go.dev/github.com/nextmv-io/sdk)

Any decision model built on Nextmv's SDK requires a [a Go module][modules]. A
module manages your dependencies, including the Nextmv SDK. Let's create one
called `tour`.

```bash
$ mkdir tour

$ cd tour

tour$ go mod init tour
go: creating new go.mod: module tour
```

Now we add Nextmv's SDK to our dependencies.

```bash
tour$ go get github.com/nextmv-io/sdk
go: added github.com/nextmv-io/sdk vX.Y.Z
```

You should now have a `go.mod` file that looks like this.

```bash
tour$ cat go.mod
module tour

go 1.18

require github.com/nextmv-io/sdk vX.Y.Z // indirect
```

Now we can create a test file that prints the SDK's version.

```bash
tour$ mkdir ehlo && cat << EOF > ehlo/main.go
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
EOF
```

We can run it using `nextmv sdk run`.

```bash
tour$ nextmv sdk run ehlo/main.go
go: downloading github.com/nextmv-io/sdk vX.Y.Z
Hello Hop vX.Y.Z
```

If you see see output like the above, you're ready to get hopping! Each of the
examples in this tour constitutes a complete `main.go`. Put them in unique
directories inside your `tour` folder and run them using the same `nextmv sdk
run` command shown above.

---

[Previous][previous] | [Next][next] | [Home][home]

[previous]: ./prerequisites.md
[next]: ../working-with-stores/creating-stores.md
[home]: ../README.md
[modules]: https://go.dev/blog/using-go-modules
