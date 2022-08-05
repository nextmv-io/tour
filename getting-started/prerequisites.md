# ![ears](../img/ears.png) Prerequisites

To get started you need Go 1.18.3. If you don't have Go installed, you can get
it [here][download]. If you already have Go, you can see which version it is by
running `go version`.

```bash
$ go version
go version go1.18.3 darwin/amd64
```

If you have a different version installed, you can install 1.18.3 [using the
`go` you already have][manage].

```bash
$ go install golang.org/dl/go1.18.3@latest

$ go1.18.3 download

$ go1.18.3 version
go version go1.18.3 darwin/amd64
```

You also need Nextmv's CLI tool. _You must have access to a private beta for
this feature. Contact [support][support] for more information_.

If not already installed and configured, follow the instructions
[here][cli-install]. The full reference for the Nextmv CLI can be found
[here][cli-reference].

Once you have successfully configured your API key, run `nextmv sdk get` to
download the necessary files. Your setup should look something like this.

```bash
$ nextmv sdk get
successfully installed sdk files
```

---

[Previous][previous] | [Next][next] | [Home][home]

[previous]: ../README.md
[next]: ./taking-the-tour.md
[home]: ../README.md
[download]: https://go.dev/dl/
[manage]:   https://go.dev/doc/manage-install
[cli-install]: https://cloud.nextmv.io/cli
[cli-reference]: https://docs.nextmv.io/cli/get-started
[support]: https://nextmv.io/contact
