// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package flydb

import (
    "encoding/json"
)

type JsonFormat struct {
}

func (this *JsonFormat) Extensions() []string {
    return []string{"json"}
}

func (this *JsonFormat) Marshal(v interface{}) ([]byte, error) {
    // TODO: make indent configurable
    return json.MarshalIndent(v, "", "    ")
}

func (this *JsonFormat) Unmarshal(b []byte) (interface{}, error) {
    var v interface{}
    if err := json.Unmarshal(b, &v); err != nil {
        return nil, err
    }

    return v, nil
}

func init() {
    RegisterFormat("json", new(JsonFormat))
}