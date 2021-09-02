# Consistent JSON Hasher

This module generates consistent json sha256 hashes irrespective of key orders.

Can be useful to

1. Compare two different JSONs
2. Generate primary keys based on JSONs and retrieve later


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
54685e7f100f9a9d85a72ec0e6d41c0b94d23391ed56cd29b94cd65cf64c9354
54685e7f100f9a9d85a72ec0e6d41c0b94d23391ed56cd29b94cd65cf64c9354
```

Irrespective of formatting on the JSON or the order of the keys, the hash will be consistent.
