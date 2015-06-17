package jsondb

import (
    "fmt"
)

type Node interface {
    SetRaw(interface{}) error
    GetRaw() interface{}
}

func CreateNodeFromRawData(rawData interface{}) (*Node, error) {
    switch typedRawData := rawData.(type) {
    case map[string]interface{}:
        return NewMapNode(typedRawData)
    case []interface{}:
        return NewArrayNode(typedRawData)
    case string, float, int, bool:
        return NewValueNode(typedRawData)
    default:
        return nil, fmt.Errorf("Invalid data type")
    }
}