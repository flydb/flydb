package flydb

import (
    "fmt"
)

type MapNode struct {
    data map[string]Node
}

func NewMapNode(v map[string]interface{}) (*MapNode, error) {
    node := &MapNode {
    }

    if err := node.SetRaw(v); err != nil {
        return nil, err
    }

    return node, nil
}

func (this *MapNode) SetRaw(raw interface{}) (error) {
    rawMap, ok := raw.(map[string]interface{})
    if !ok {
        return fmt.Errorf("raw data is not a map")
    }

    data := make(map[string]Node)
    for k, v := range rawMap {
        node, err := CreateNodeFromRawData(v)
        if err != nil {
            return err
        }
        data[k] = node
    }

    this.data = data
    return nil
}

func (this *MapNode) GetRaw() interface{} {
    result := make(map[string]interface{})
    for k, v := range this.data {
        result[k] = v.GetRaw()
    }

    return result
}

func (this *MapNode) Get(k string) (Node, bool) {
    node, ok := this.data[k]
    return node, ok
}

func (this *MapNode) Set(k string, node Node) {
    this.data[k] = node
}

func (this *MapNode) Has(k string) bool {
    _, ok := this.data[k]
    return ok
}

func (this *MapNode) Delete(k string) {
    delete(this.data, k)
}