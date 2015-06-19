// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package flydb

import (
    "testing"
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

func GetTestDB() *Database {
    db, err := Open("tests/db.json")
    if err != nil {
        panic(err)
    }

    return db
}