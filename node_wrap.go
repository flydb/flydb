// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package flydb

import (
    "fmt"
)

type WrapNode interface {
    SetRaw(interface{}) error
    GetRaw() interface{}
}

func CreateWrapNodeFromRawData(rawData interface{}) (WrapNode, error) {
    switch typedRawData := rawData.(type) {
    case map[string]interface{}:
        return NewMapNode(typedRawData)
    case []interface{}:
        return NewArrayNode(typedRawData)
    case string, float32, float64, int, int64, bool:
        return NewValueNode(rawData)
    default:
        return nil, fmt.Errorf("Invalid data type: %v", rawData)
    }
}

func CreateWrapNode(data interface{}) (WrapNode, error) {
    switch typedData := data.(type) {
    case WrapNode:
        return typedData, nil
    default:
        return CreateWrapNodeFromRawData(data)
    }
}