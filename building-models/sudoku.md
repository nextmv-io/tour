# ![ears](../img/ears.png) Sudoku

🏗️ The text for this section still needs to be written. You can still try out
the [source listing][source], though.

```bash
go run -trimpath building-models/sudoku/main.go \
    -hop.runner.input.path building-models/sudoku/input.json \
    -hop.solver.limits.solutions 1 |\
jq -c .store
```

```json
[
    [8,2,9,4,5,6,7,3,1],
    [1,5,7,3,8,2,6,9,4],
    [4,3,6,9,1,7,5,8,2],
    [6,9,3,5,2,8,4,1,7],
    [2,4,5,1,7,3,9,6,8],
    [7,8,1,6,4,9,3,2,5],
    [3,1,8,7,6,5,2,4,9],
    [5,7,2,8,9,4,1,4,3],
    [9,6,4,2,3,1,8,5,7]
]
```

[source]: sudoku/main.go
