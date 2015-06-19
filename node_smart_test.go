package flydb

import (
    "testing"
)

func TestGet(t *testing.T) {
    db := GetTestDB()
    root := db.MustGet(nil)
    root.MustGet("users")
}