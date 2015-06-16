package jsondb

import (
    "fmt"
)

// Array node contains list of nodes
type ArrayNode struct {
    data []*Node
}

// Convert node to JSON bytes
func (this *ArrayNode) Marshal() ([]byte, error) {

}

func (this *ArrayNode) Unmarshal(b []byte) error {

}

func (this *ArrayNode) Append(node *Node) {
    this.data = append(this.data, node)
}

func (this *ArrayNode) Get(i int) (*Node, error) {
    if i < 0 || i >= len(this.data) {
        return nil, fmt.Errorf("key out of range: %d", i)
    }

    return this.data[i], nil
}

func (this *ArrayNode) Set(i int, node *Node) error {
    if i < 0 || i >= len(this.data) {
        return nil, fmt.Errorf("key out of range: %d", i)
    }

    this.data[i] = node
    return nil
}

func (this *ArrayNode) Delete(key int) {
    if key < 0 || key >= len(this.data) {
        return
    }

    this.data = append(this.data[0:key], this.data[key+1:])
}

func (this *ArrayNode) Length() int {
    return len(this.data)
}