# flydb

    In progress...

Pure Go database with simple format.

## Features

- Support different data file format: `JSON`, `YAML`, `XML`
- Powerful access API

## Limits

- Performance: flydb is not built for performance, so you may not want to use it with huge data

## Usage

### Embed in Golang

```go
package main

import (
    "github.com/flydb/flydb"
    "log"
)

func main() {
    db, err := flydb.Open("/path/to/db.json")
    if err != nil {
        fmt.Errorf("cannot open database")
    }
    email := db.MustGet("users.3.email").MustVallue().MustString()
    log.Println(email)
}
```

### Server

TODO

## Data Format

Currently support JSON, YAML and XML, you can convert between all these formats.
