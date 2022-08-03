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

You also need Nextmv's cli tool. _You must have access to a private beta for
this feature. Contact [support][support] for more information_.

If not already installed and configured, follow the instructions
[here][cli-install]. The full reference for the Nextmv cli can be found
[here][cli-reference].

Once you have successfully configured your api key, run `nextmv sdk get` to
download the necessary binary files to `~/.nextmv/lib`. Add a
`NEXTMV_LIBRARY_PATH` environment variable pointing to them. Your setup should
look something like this.

```bash
$ nextmv sdk get
successfully installed sdk files

$ ls ~/.nextmv/lib
nextmv-run-cli-v0.16.0-dev.0-4-go1.18.3-darwin-amd64.so        
nextmv-sdk-v0.16.0-dev.0-4-go1.18.3-darwin-amd64.so
nextmv-run-http-v0.16.0-dev.0-4-go1.18.3-darwin-amd64.so

$ export NEXTMV_LIBRARY_PATH=~/.nextmv/lib

$ echo $NEXTMV_LIBRARY_PATH
~/.nextmv/lib
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
