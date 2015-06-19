// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package flydb

import (
    "fmt"
    "gopkg.in/yaml.v2"
)

type YamlFormat struct {
}

func (this *YamlFormat) Extensions() []string {
    return []string{"yaml", "yml"}
}

func (this *YamlFormat) Marshal(v interface{}) ([]byte, error) {
    return yaml.Marshal(v)
}

func (this *YamlFormat) Unmarshal(b []byte) (interface{}, error) {
    var v interface{}
    if err := yaml.Unmarshal(b, &v); err != nil {
        return nil, err
    }

    return fixYamlMap(v)
}

// Fix map, note that the key type of YAML map is interface{}, we have to 
// convert it to string
func fixYamlMap(data interface{}) (interface{}, error) {
    switch typedData := data.(type) {
    case map[interface{}]interface{}:
        result := make(map[string]interface{})
        for k, v := range typedData {
            stringKey, ok := k.(string)
            if !ok {
                return nil, fmt.Errorf("unknown key type")
            }
            child, err := fixYamlMap(v)
            if err != nil {
                return nil, err
            }

            result[stringKey] = child
        }

        return result, nil
    case []interface{}:
        for k, v := range typedData {
            child, err := fixYamlMap(v)
            if err != nil {
                return nil, err
            }

            typedData[k] = child
        }

        return typedData, nil
    default:
        return data, nil
    }
}

func init() {
    RegisterFormat("yaml", new(YamlFormat))
}