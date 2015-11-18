
# cacheddns

[![GoDoc](https://godoc.org/github.com/ernestoalejo/cacheddns?status.svg)](https://godoc.org/github.com/ernestoalejo/cacheddns)

Get the IP of a domain, caching the result temporarily in memory.


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

