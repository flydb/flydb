package main

import "github.com/selphish/jsondb"

func main() {
    db, _ := jsondb.Open("db.json")

    // get
    println db.MustGetString("users.0.username")

    // set
    db.Set("users.0.star", 100)

    // filter
    db.MustGet("users").Filter(func(n *db.Node) bool {
        return true
    })

}