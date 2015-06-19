// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package flydb

import (
    "testing"
    "encoding/json"
)

func doTestFormat(t *testing.T, format Format) {
    jsonData := `
{
    "users": [
        {
            "username": "user1",
            "email": "email1@test.com"
        },
        {
            "username": "user2",
            "email": "email2@test.com"
        },
        {
            "username": "user3",
            "email": "email3@test.com"
        }
    ]
}
    `
    var data interface{}
    if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
        panic(err)
    }

    bytes, err := format.Marshal(data)
    if err != nil {
        t.Fatal(err)
    }
    data, err = format.Unmarshal(bytes)
    if err != nil {
        t.Fatal(err)
    }

    node, err := CreateNode(data)
    if err != nil {
        t.Fatal(err)
    }

    if node.MustGet("users.1.username").MustValue().MustString() != "user2" {
        t.Fail()
    }
}