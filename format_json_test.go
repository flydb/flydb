// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package flydb

import (
    "testing"
)

func TestFormatJson(t *testing.T) {
    format := GetFormat("json")

    doTestFormat(t, format)
}