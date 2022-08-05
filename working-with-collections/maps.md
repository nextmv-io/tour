# ![ears](../img/ears.png) Maps

Just like stores provide an immutable slice type and various methods for
creating new slices from existing ones, it also provides a map collection. Like
slices, maps can store any type of value. However, they only allow either `int`
or `string` keys.

A map is initialized empty and therefore requires its key and value types.

```go
s1 := store.New()
x := store.NewMap[string, float64](s1) // map[string]float64{}
```

We can assign values to keys in a map using its `Set` method. Like other types
associated with a store, instead of mutating the underlying data in a maps, this
returns a change to apply to a new store.

```go
s2 := s1.Apply( // map[string]float64{"pi": 3.14, "e": 2.72}
    x.Set("pi", 3.14),
    x.Set("e", 2.72),
)
```

Maps can assign new values to existing keys through subsequent calls to `Set`.
They can also remove keys entirely using their `Delete` method.

```go
s3 := s2.Apply(x.Delete("e")) // map[string]float64{"pi": 3.14}
```

Among the other methods available on the map collection, the `Map` method
returns its underlying representation.

```go
fmt.Println(x.Map(s3))
```

## Exercises

* Try to guess what `s1`, `s2`, and `s3` contain. Run the [source][source] and
  see if you are right.
* Create a map with `int` keys and values of a custom type. Set values on the
  map and retrieve its underlying representation.

---

[Previous][previous] | [Next][next] | [Home][home]

[previous]: ./slices.md
[next]: ./domains.md
[home]: ../README.md
[source]: maps/main.go
