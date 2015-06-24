// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package flydb

import (
    "testing"
    "time"
)

func TestOpen(t *testing.T) {
    _, err := Open("tests/db.json")
    if err != nil {
        t.Error(err)
    }
}

func TestSave(t *testing.T) {
    db := GetTestDB()
    if err := db.SaveAs("tests/tmp/saveas.json"); err != nil {
        t.Error(err)
    }
}

func TestAutoSave(t *testing.T) {
    db := New(Config {
        Path: "tests/autosave.json",
        Save: true,
        SaveInterval: 1000,
    })
    err := db.Open()
    if err != nil {
        t.Fatal(err)
    }

    counter := db.Root().MustGet("counter")
    v := counter.MustValue().MustInt64()
    counter.Set(".", v + 1)
    time.Sleep(1500*time.Millisecond)
    db2, _ := Open("tests/autosave.json")
    if db2.Root().MustGet("counter").MustValue().MustInt64() != v + 1 {
        t.Fatal()
    }
}

func GetTestDB() *Database {
    db, err := Open("tests/db.json")
    if err != nil {
        panic(err)
    }

    return db
}