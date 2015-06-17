package flydb

import (
    "testing"
)

func TestOpen(t *testing.T) {
    db, err := Open("test.json")
    if err != nil {
        t.Error(err)
    }

    t.Log(db.Get("users.0.username").Value().String())
}