
# cacheddns

Resolve domains through a DNS caching the results in memory.

[![GoDoc](https://godoc.org/github.com/ernestoalejo/cacheddns?status.svg)](https://godoc.org/github.com/ernestoalejo/cacheddns)

Go HTTP handlers using context.


## Installation

```shell
go get github.com/ernestoalejo/cacheddns
```


### Usage

```go
package main

import (
  "log"
  "time"

  "github.com/ernestoalejo/cacheddns"
)

func main() {
  myDomain := cacheddns.New("www.google.com", 30*time.Second)

  // Will resolve the domain once every 30 seconds
  address, err := myDomain.Resolve()
  if err != nil {
    log.Fatal(err)
  }
}
```

