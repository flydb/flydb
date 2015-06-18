package flydb

import (
    "testing"
)

func TestOpen(t *testing.T) {
    db, err := Open("tests/db.json")
    if err != nil {
        t.Error(err)
    }

    t.Log(db.MustGet("users.0.username").MustValue().String())
    t.Log(db.MustGet("users.1.email").MustValue().String())

    t.Log(db.SaveAs("tests/tmp/saveas.json"))
}