package main

import "github.com/flydb/flydb"

func main() {
    db, _ := flydb.Open("db.json")

    // get
    println(db.Root().MustGet("users.0.username").MustValue().MustString())

    // set
    db.Root().Set("users.0.star", 100)

    println(db.Root().MustGet("users.0.star").MustValue().MustInt())

    // filter
    // db.MustGet("users").Filter(func(n *db.Node) bool {
    //     return true
    // })

}