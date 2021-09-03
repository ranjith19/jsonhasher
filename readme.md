# Consistent JSON Hasher in go

This module generates consistent json sha256 hashes irrespective of key orders.

Use cases:

1. Compare two different JSONs for equality
2. Generate keys based on JSON values in key value stores. (for example caching API response combination of inputs.)


## Example Usage



```go
package main

import (
    "fmt"
    "github.com/ranjith19/jsonhasher"
)

func main() {
    input1 := `{"a": 1, "c": {"x": true, "y": false}, "b": 2}`
    input2 := `{"b": 2, 
    "a": 1, "c": {"x": true, "y": false}}`

    h1, _ := jsonhasher.HashJsonString(input1)
    h2, _ := jsonhasher.HashJsonString(input2)

    fmt.Println(*h1)
    fmt.Println(*h2)
}
```

Will output

```txt
4af2b0287001172868b69c6d7531661e6f0d6ec0f4f506876c387d0b4b7ecbdf
4af2b0287001172868b69c6d7531661e6f0d6ec0f4f506876c387d0b4b7ecbdf
```

Irrespective of formatting on the JSON or the order of the keys, the hash will be consistent.

Try this at [go playground](https://play.golang.org/p/awmKpeeeqvS). Another example [here](https://play.golang.org/p/Av9jvDo5xap)


## Longer/Shorter hashes

You can use `HashJsonStringSha1` or `HashJsonStringSha512` methods.

## testing

```
go test
go test -covermode=count -coverprofile=coverage.out
go tool cover -html=coverage.out
```