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
    return json.Marshal(v)
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