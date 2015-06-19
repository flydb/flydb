// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package flydb

import (
    "testing"
)

func TestGet(t *testing.T) {
    db := GetTestDB()
    root := db.Root()
    if root.MustGet("users.0.email").MustValue().MustString() != "email1@test.com" {
        t.Fail()
    }
}

func TestSet(t *testing.T) {
    db := GetTestDB()
    root := db.Root()
    err := root.Set("x.y.z", "abc")
    if err != nil {
        t.Error(err)
    }
    if root.MustGet("x.y.z").MustValue().MustString() != "abc" {
        t.Fail()
    }
}