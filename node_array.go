package flydb

import (
    "fmt"
)

func NewArrayNode(v []interface{}) (*ArrayNode, error) {
    node := &ArrayNode {
    }

    if err := node.SetRaw(v); err != nil {
        return nil, err
    }

    return node, nil
}

// Array node contains list of nodes
type ArrayNode struct {
    data []Node
}

func (this *ArrayNode) SetRaw(raw interface{}) (error) {
    rawArray, ok := raw.([]interface{})
    if !ok {
        return fmt.Errorf("raw data is not an array")
    }
    data := make([]Node, len(rawArray))
    for k, v := range rawArray {
        node, err := CreateNodeFromRawData(v)
        if err != nil {
            return err
        }

        data[k] = node
    }

    this.data = data
    return nil
}

func (this *ArrayNode) GetRaw() interface{} {
    result := make([]interface{}, len(this.data))
    for i, v := range this.data {
        result[i] = v.GetRaw()
    }

    return result
}

func (this *ArrayNode) Append(node Node) {
    this.data = append(this.data, node)
}

func (this *ArrayNode) Get(i int) (Node, error) {
    if i < 0 || i >= len(this.data) {
        return nil, fmt.Errorf("key out of range: %d", i)
    }

    return this.data[i], nil
}

func (this *ArrayNode) Set(i int, node Node) error {
    if i < 0 || i >= len(this.data) {
        return fmt.Errorf("key out of range: %d", i)
    }

    this.data[i] = node
    return nil
}

func (this *ArrayNode) Delete(key int) {
    if key < 0 || key >= len(this.data) {
        return
    }

    this.data = append(this.data[0:key], this.data[key+1:]...)
}

func (this *ArrayNode) Length() int {
    return len(this.data)
}