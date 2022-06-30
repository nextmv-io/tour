# ![ears](../img/ears.png) Input and Output

We're ready to start building models that actually do something. But first, we
need to take a quick detour to understand how Hop runners let us read input data
into a model and output formatted JSON data to make decisions.

A runner is responsible for reading input data, setting up the solver execution
environment, and writing output to a desired location. Runners make it easy to
switch between different environments which may need to read data from different
places or handle timeouts differently. They are the key to writing model code
locally with confidence that the model with behave the same in production.

The code below creates and empty store and passes it to a runner. Technically
this is a model, though it doesn't anything. What it does do is let us see the
whole process of reading data, building a model, and solving that model.

```go
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
```

Most Hop models begin with a call to `run.Run` in a `main` function. `run.Run`
requires a handler. A handler reads data _of any JSON-unmarshalable type_, pulls
solver options out of the environment or command line, and constructs a solver.
The runner unmarshals the input for your and knows what to do with the solver.

In our handler here, we don't care what type the input data is, so we label it
as `any`. Usually we would use something like `n int` or `x foo`, where `foo`
can be any structure that follows [Go's JSON decoding rules][json].

Let's run this [empty model][source] and pipe the output into [`jq`][jq].

```bash
echo 42 | go run -trimpath building-models/input-and-output/main.go | jq
```

You should see output similar to the following. It hows how Hop was configured,
statistics about its search, and so on. The `store` field is `null` because our
model doesn't do anything, but that will change in the next sections.

```json
{
  "version": {
    "sdk": "v0.16.0-dev.0-3"
  },
  "options": {
    "diagram": {
      "expansion": {
        "limit": 0
      },
      "width": 10
    },
    "limits": {
      "duration": "10s"
    },
    "search": {
      "buffer": 100
    },
    "sense": "satisfy"
  },
  "store": null,
  "statistics": {
    "search": {
      "generated": 0,
      "filtered": 0,
      "expanded": 0,
      "reduced": 0,
      "restricted": 0,
      "deferred": 0,
      "explored": 1,
      "solutions": 0
    },
    "time": {
      "elapsed": "297.678µs",
      "elapsed_seconds": 0.000297678,
      "start": "2022-06-30T13:45:03.280816884-04:00"
    }
  }
}
```

The runner also accepts a wealth of command line flags and environment
variables. You can see these by passing the `-h` flag.

```bash
go run -trimpath building-models/input-and-output/main.go -h
```

Some flags and variables are specific to the runner. By default, Hop uses the
CLI runner, though runners for things like HTTP, Lambda, and S3 triggered
Lambda are also available. You can read more about this in our [docs][docs].

## Exercises

* Run the model with `-h`. What are the most useful flags to you?
* Try setting the solutions output flag to `all` and `last`. How does this
  change the resulting JSON?

[source]: input-and-output/main.go
[json]:   https://pkg.go.dev/encoding/json
[docs]:   https://docs.nextmv.io/overview/decision-stack/runners
