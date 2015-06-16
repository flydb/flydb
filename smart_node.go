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

func (this *SmartNode) Get(path interface{}) (*SmartNode, error) {
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

func (this *SmartNode) Set(path interface{}, v interface{}) (error) {
    steps := parsePath(path)
    var node Node;
    switch typedValue := v.(type) {
        case Node {

        }
    }
    if !steps {
    }
}

func (this *SmartNode) Delete(path interface{}) error {
    steps := parsePath(path)
    if !steps {
        return fmt.Errorf("cannot delete root")
    }

    if len(steps) == 1 {
        switch node := this.node.(type) {
        case *MapNode:
            node.Delete(steps[0])
            return nil
        case *ArrayNode:
            var key, ok := strconv.Atoi(steps[0])
            if !ok {
                return fmt.Errorf("key of array is not int")
            }
            return node.Delete(key)
        default:
            return fmt.Errorf("value node does not have children")
        }
    }

    parent, err := this.Get(steps[0:len(steps) - 1])
    if err != nil {
        return err
    }

    return parent.Delete(steps[1:])
}

func (this *SmartNode) Array() (*ArrayNode, error) {
    node, ok := this.node.(*ArrayNode)
    if !ok {
        return nil, fmt.Errorf("not an array node")
    }

    return node, nil
}

func (this *SmartNode) Map() (*MapNode, error) {
    node, ok := this.node.(*MapNode)
    if !ok {
        return nil, fmt.Errorf("not a map node")
    }

    return node, nil
}

func (this *SmartNode) Value() (*ValueNode, error) {
    node, ok := this.node.(*ValueNode)
    if !ok {
        return nil, fmt.Errorf("not a value node")
    }

    return node, nil
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