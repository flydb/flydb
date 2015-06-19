// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package flydb

import (
    "fmt"
    "strconv"
)

func CreateNode(data interface{}) (*Node, error) {
    switch typedData := data.(type) {
    case *Node:
        return typedData, nil
    case WrapNode:
        return &Node{typedData}, nil
    default:
        node, err := CreateWrapNodeFromRawData(data)
        if err != nil {
            return nil, err
        }

        return &Node{node}, nil
    }
}

// Wrapper for all types of node
type Node struct {
    node WrapNode
}

func (this *Node) GetRaw() interface{} {
    return this.node.GetRaw()
}

// Get node with path
func (this *Node) Get(path interface{}) (*Node, error) {
    steps := parsePath(path)
    if len(steps) == 0 {
        return this, nil
    }

    switch node := this.node.(type) {
        case *MapNode:
            next, ok := node.Get(steps[0])
            if !ok {
                return nil, fmt.Errorf("path does not exist: %s", steps[0])
            }

            return next.Get(steps[1:])
        case *ArrayNode:
            key, err := strconv.Atoi(steps[0])
            if err != nil {
                return nil, fmt.Errorf("key for array node should be int: %s", steps[0])
            }

            next, err := node.Get(key)
            if err != nil {
                return nil, err
            }

            return next.Get(steps[1:])
        default:
            return nil, fmt.Errorf("cannot access value node with key: %v", steps)
    }
}

func (this *Node) MustGet(path interface{}) (*Node) {
    n, err := this.Get(path)
    if err != nil {
        panic(err)
    }

    return n
}

// Set node value by path
func (this *Node) Set(path interface{}, v interface{}) (error) {
    steps := parsePath(path)

    // Set value for current node
    if len(steps) == 0 {
        node, err := CreateWrapNode(v)
        if err != nil {
            return err
        }

        this.node = node
        return nil
    }

    firstStep := steps[0]
    switch node := this.node.(type) {
    case *MapNode:
        if !node.Has(firstStep) {
            node.Set(firstStep, make(map[string]interface{}))
        }
        child, ok := node.Get(firstStep)
        if !ok {
            return fmt.Errorf("cannot create empty map")
        }

        err := child.Set(steps[1:], v)
        return err
    case *ArrayNode:
        key, err := strconv.Atoi(firstStep)
        if err != nil {
            return fmt.Errorf("invalid key for array node")
        }

        child, err := node.Get(key)
        if err != nil {
            return err
        }

        return child.Set(steps[1:], v)
    default:
        return fmt.Errorf("value node does not have children")
    }
}

// Delete node by path
func (this *Node) Delete(path interface{}) error {
    steps := parsePath(path)
    if len(steps) == 0 {
        return fmt.Errorf("cannot delete root")
    }

    if len(steps) == 1 {
        switch node := this.node.(type) {
        case *MapNode:
            node.Delete(steps[0])
            return nil
        case *ArrayNode:
            key, err := strconv.Atoi(steps[0])
            if err != nil {
                return fmt.Errorf("key of array is not int")
            }
            node.Delete(key)
            return nil
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
func (this *Node) Has(path interface{}) bool {
    _, err := this.Get(path)
    if err != nil {
        return false
    }

    return true
}

// Convert to array node
func (this *Node) Array() (*ArrayNode, error) {
    node, ok := this.node.(*ArrayNode)
    if !ok {
        return nil, fmt.Errorf("not an array node")
    }

    return node, nil
}

func (this *Node) MustArray() (*ArrayNode) {
    n, err := this.Array()
    if err != nil {
        panic(err)
    }

    return n
}

// Convert to map node
func (this *Node) Map() (*MapNode, error) {
    node, ok := this.node.(*MapNode)
    if !ok {
        return nil, fmt.Errorf("not a map node")
    }

    return node, nil
}

func (this *Node) MustMap() (*MapNode) {
    n, err := this.Map()
    if err != nil {
        panic(err)
    }

    return n
}

// Convert to value node
func (this *Node) Value() (*ValueNode, error) {
    node, ok := this.node.(*ValueNode)
    if !ok {
        return nil, fmt.Errorf("not a value node")
    }

    return node, nil
}

func (this *Node) MustValue() (*ValueNode) {
    n, err := this.Value()
    if err != nil {
        panic(err)
    }

    return n
}

// Check if this is an array node
func (this *Node) IsArray() bool {
    _, ok := this.node.(*ArrayNode)
    return ok
}

// Check if this is a map node
func (this *Node) IsMap() bool {
    _, ok := this.node.(*MapNode)
    return ok
}

// Check if this is a value node
func (this *Node) IsValue() bool {
    _, ok := this.node.(*ValueNode)
    return ok
}