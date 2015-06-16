package jsondb

import (
    "fmt"
    "strconv"
)

// Wrapper for all types of node
type SmartNode struct {
    node *Node
}

// Convert node to json bytes
func (this *SmartNode) Marshal() ([]byte, error) {
    return this.node.Marshal()
}

// Convert json bytes to node
func (this *SmartNode) Unmarshal(b []byte) error {
    return this.node.Unmarshal(b)
}

// Get node with path
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
                next,
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
                next,
            }

            return nextNode.Get(steps[1:])
        default:
            return nil, fmt.Errorf("cannot access value node with key")
    }
}

// Set node value by path
func (this *SmartNode) Set(path interface{}, v interface{}) (error) {
    steps := parsePath(path)

    // Set value for current node
    if !steps {
        if node, ok := v.(*Node); ok {
            this.node = node
            return nil
        }

        if node, ok := this.node.(*ValueNode); ok {
            return node.SetValue(v)
        }

        return fmt.Errorf("set by path failed")
    }

    firstStep := steps[0]
    switch node := this.node.(type) {
    case *MapNode:
        if !node.Has(firstStep) {
            node.Set(firstStep, NewMapNode())
        }
        child, err := node.Get(firstStep)
        if err != nil {
            return err
        }
        return &SmartNode{
            child,
        }.Set(steps[1:], v)
    case *ArrayNode:
        key, ok := strconv.Atoi(firstStep)
        if !ok {
            return fmt.Errorf("invalid key for array node")
        }

        child, err := node.Get(key)
        if err != nil {
            return err
        }

        return &SmartNode{
            child,
        }.Set(steps[1:], v)
    default:
        return fmt.Errorf("value node does not have children")
    }
}

// Delete node by path
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
            key, ok := strconv.Atoi(steps[0])
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

// Check path existance
func (this *SmartNode) Has(path interface{}) bool {
    _, err := this.Get(path)
    if err != nil {
        return false
    }

    return true
}

// Convert to array node
func (this *SmartNode) Array() (*ArrayNode, error) {
    node, ok := this.node.(*ArrayNode)
    if !ok {
        return nil, fmt.Errorf("not an array node")
    }

    return node, nil
}

// Convert to map node
func (this *SmartNode) Map() (*MapNode, error) {
    node, ok := this.node.(*MapNode)
    if !ok {
        return nil, fmt.Errorf("not a map node")
    }

    return node, nil
}

// Convert to value node
func (this *SmartNode) Value() (*ValueNode, error) {
    node, ok := this.node.(*ValueNode)
    if !ok {
        return nil, fmt.Errorf("not a value node")
    }

    return node, nil
}

// Check if this is an array node
func (this *SmartNode) IsArray() bool {
    _, ok := this.node.(*ArrayNode)
    return ok
}

// Check if this is a map node
func (this *SmartNode) IsMap() bool {
    _, ok := this.node.(*MapNode)
    return ok
}

// Check if this is a value node
func (this *SmartNode) IsValue() bool {
    _, ok := this.node.(*ValueNode)
    return ok
}