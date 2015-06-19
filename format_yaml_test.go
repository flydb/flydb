// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package flydb

import (
    "testing"
)

func TestFormatYaml(t *testing.T) {
    format := GetFormat("yaml")

    doTestFormat(t, format)
}