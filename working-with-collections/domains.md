# ![ears](../img/ears.png) Domains

Domains are a special type. A domain stores integers which typically represent
potential choices. For example, a domain may represent the hours a shift might
start or the destinations a traveler could arrive at.

Structurally, a domain is an ordered, compact, set of integers. Domains maintain
a minimal representation of ranges as we apply operators to them to create new
ones. They are also, conveniently, immutable.

Nextmv's SDK provides two domain types: `model.Domain` and `store.Domain`.
`model.Domain` is just an integer domain, unattached to a store. `store.Domain`
has many of the same methods but with similar mechanics to `store.Slice` and
`store.Map`. Thus `model.Domain` is the underlying type for `store.Domain`,
which must be associated with a store.

Let's create a few domains to see how they work.

```go
s1 := store.New()
d1 := store.NewDomain(s1)
d2 := store.NewDomain(s1, model.NewRange(1, 10))
d3 := store.NewDomain(s1, model.NewRange(-5, 5), model.NewRange(15, 25))
```

We can pass as many integer ranges as we want to `store.NewDomain`. If we pass
none, the domain is empty, like `d1`. `d2` contains the integers 1 through 10,
while `d3` contains two disjoint ranges of integers, -5 though 5 and 15 through
25.

Let's apply a change set to these domains and encode the store into JSON.

```go
s2 := s1.Apply(
    d1.Add(42, 43, 44),
    d2.Remove(2, 4, 6, 8, 10),
    d3.AtLeast(10),
)

enc := json.NewEncoder(os.Stdout)
enc.Encode(s2)
```

In `s2`, `d1` contains 42 through 44, `d2` contains only odd numbers, and `d3`
contains the range 15 through 25. Note that each of these domains is encoded in
a compact JSON representation. This is similar to its internal data.

```json
[[[42,44]],[1,3,5,7,9],[[15,25]]]
```

Frequently, we want to create and operate on multiple domains at once. We can
use functions like `store.NewDomain` or `store.Repeat` to create a slice of
related domains. In the code below, we create five domains, each containing the
values 1 through 10.

```go
s3 := store.New()
d := store.Repeat(s3, 5, model.NewDomain(model.NewRange(1, 10)))
```

The `store.Domains` type has many of the same methods as the `store.Domain`
type. Change methods require a domain index to modify as their first argument.

```go
s4 := s3.Apply(
    d.Add(0, 42),
    d.Assign(2, 5),
    d.Remove(4, 2, 3, 4),
)
enc.Encode(s4)
```

The change set above results in the following domains.

```json
[[[[1,10],42],[[1,10]],5,[[1,10]],[1,[5,10]]]]
```

There are many methods available on domains. Some allow us to modify them, while
others help us select an individual domain from a collection of them. Take a
look at the Go package documentation to see what domains have to offer. Run the
[source][source] to get the outputs described here.

## Exercises

* Create a domain on a store and a domain unattached to a store. Modify both of
  these domains in various ways.
* Create multiple distinct domains on a store. Use a selector method, like
  `Smallest` or `Largest` to select an individual domain by index. Assign that
  domain a value.
* Create a collection of domains each containing more than one value. Remove
  values from the domains until calling `Singleton` returns true.

---

[Previous][previous] | [Next][next] | [Home][home]

[previous]: ./maps.md
[next]: ../building-models/input-and-output.md
[home]: ../README.md
[source]: domains/main.go
