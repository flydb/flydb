// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package flydb

import (
    "encoding/xml"
)

type XMLFormat struct {
}

func (this *XMLFormat) Extensions() []string {
    return []string{"xml"}
}

func (this *XMLFormat) Marshal(v interface{}) ([]byte, error) {
    // TODO: make indent configurable
    return xml.MarshalIndent(v, "", "    ")
}

func (this *XMLFormat) Unmarshal(b []byte) (interface{}, error) {
    var v interface{}
    if err := xml.Unmarshal(b, &v); err != nil {
        return nil, err
    }

    return v, nil
}

func init() {
    RegisterFormat("xml", new(XMLFormat))
}