# ![ears](../img/ears.png) Searching

The store's `Generate` function is possibly the most powerful one, as it allows
you to define the guardrails to generate new stores from existing ones. This
creates a search tree that the solver uses to find the best operationally valid
store. A store is operationally valid if all decisions have been made and they
are valid.

We are going to search for all permutations of the positive integer numbers
that go up to a specific number, e.g.:

 ```txt
1 -> [
  [1]
]
2 -> [
  [1,2],
  [2,1]
]
3 -> [
  [1,2,3]
  [1,3,2]
  [2,1,3]
  [2,3,1]
  [3,1,2]
  [3,2,1]  
]
 ```

You can see that as we increase the number, this search becomes non-trivial.
Starting with an integer input `n`, define the root store and a domain of the
unused integers:

```go
root := store.New()
unused := store.NewDomain(root, model.NewRange(1, n))
```

We are going to use a slice to store the permutations as we search for them.

```go
permutation := store.NewSlice[int](root)
```

Our store is operationally valid if we have found _all_ permutations, i.e.: the
`unused` domain is empty.

```go
root = root.Validate(unused.Empty)
```

For a simple output format, we can visualize the `permutation` slice.

```go
root = root.Format(func(s store.Store) any { return permutation.Slice(s) })
```

From an existing parent store, we must define the rules to generate child
stores. We can do so through an `Eager` or `Lazy` generator. In the following
code snippet, we lazily generate new children stores by getting the values that
haven't been used for a parent store. For each unused value, we append it to
our permutations and remove it from the unused domain.

```go
root = root.Generate(func(s store.Store) store.Generator {
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
```

The lazy generator will generate new stores when the solver needs them. On the
other hand, an eager generator will generate all children stores upfront. Given
that we only care about finding all permutations, it is sufficient to satisfy
operational validity by summoning a `Satisfier`, which is a type of solver.

```go
func handler(n int, opt store.Options) (store.Solver, error) {
    // code goes here
    return root.Satisfier(opt), nil
}
```

When seeking to maximize or minimize a value, you can use the `Value` function
on the store to define its value and the `Maximizer` or `Minimizer` solver,
respectively.

For `n = 3`, you can run the following command and [source][source] to observe
the corresponding result.

```bash
$ echo 3 |\
  nextmv sdk run \
    building-models/searching/main.go \
    -hop.runner.output.stream |\
  jq -c .store
```

```json
[1,2,3]
[1,3,2]
[2,1,3]
[2,3,1]
[3,1,2]
[3,2,1]
```

Use the Nextmv cli to look at complete examples of more advanced search
problems:

```bash
$ nextmv sdk init -t knapsack
Successfully generated the knapsack template.
$ nextmv sdk init -t sudoku
Successfully generated the sudoku template.
$ nextmv sdk init -t shift-scheduling
Successfully generated the shift-scheduling template.
```

## Exercises

* Run the tutorial for different values of `n`, such as 4.
* Implement this tutorial using an `Eager` generator.
* Use `Value` and `Minimizer` to look for the permutation that has the smallest
  absolute distance between its numbers, i.e.: the permutation `[1,3,4,2]` has
  a distance of `|3-1|+|4-3]+|2-4| = 2+1+2 = 5`. On the other hand, the
  permutation `[1,2,3,4]` has a distance of 3.
* Use `Value` and `Maximizer` to look for the permutation that has the largest
  absolute distance between its numbers.

---

[Previous][previous] | [Home][home]

[previous]: ./input-and-output.md
[home]: ../README.md
[source]: searching/main.go
