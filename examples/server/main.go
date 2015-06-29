package main

import (
    "github.com/flydb/flydb"
    "fmt"
)

func main() {
    const LISTEN = "127.0.0.1:8080"
    fmt.Println("Starting server, listen " + LISTEN)
    flydb.Serve("examples/db.json", LISTEN)
}