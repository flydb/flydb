package jsondb

import (
    "fmt"
    "strconv"
)

type SmartNode struct {
    node *Node
}

func (this *SmartNode) Marshal() ([]byte, error) {
    return this.node.Marshal()
}

func (this *SmartNode) Unmarshal(b []byte) error {
    return this.node.Unmarshal(b)
}

func (this *SmartNode) Get(path) (*SmartNode, error) {
    steps := parsePath(path)
    if !steps {
        return this, nil
    }

    switch node := this.node.(type) {
        case *MapNode:
            next, ok := node.Get(steps[0])
            if !ok {
                return nil, fmt.Errorf("path does not exist: %s", steps[0])
            }

            nextNode = &SmartNode {
                next
            }

            return nextNode.Get(steps[1:])
        case *ArrayNode:
            key, err := strings.Atoi(steps[0])
            if err {
                return nil, fmt.Errorf("key for array node should be int: %s", steps[0])
            }

            next, ok := node.Get(key)
            if !ok {
                return nil, fmt.Errorf("path does not exist: %d", key)
            }

            nextNode = &SmartNode {
                next
            }

            return nextNode.Get(steps[1:])
        default:
            return nil, fmt.Errorf("cannot access value node with key")
    }
}

func (this *SmartNode) Set(path, node interface{}) (error) {
    steps := parsePath(path)

}

func (this *SmartNode) Delete(path) error {

}

func (this *SmartNode) Array() (*ArrayNode, error) {

}

func (this *SmartNode) Map() (*MapNode, error) {

}

func (this *SmartNode) Value() (*ValueNode, error) {

}

func (this *SmartNode) IsArray() bool {
    _, ok := this.node.(*ArrayNode)
    return ok
}

func (this *SmartNode) IsMap() bool {
    _, ok := this.node.(*MapNode)
    return ok
}

func (this *SmartNode) IsValue() bool {
    _, ok := this.node.(*ValueNode)
    return ok
}