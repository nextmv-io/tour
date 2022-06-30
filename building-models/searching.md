# ![ears](../img/ears.png) Searching

🏗️ The text for this section still needs to be written. You can still try out
the [source listing][source], though.

```bash
echo 3 |\
  go run \
    -trimpath \
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

[source]: searching/main.go
