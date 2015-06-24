// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package flydb

import (
    "fmt"
)

func NewValueNode(v interface{}) (*ValueNode, error) {
    node := & ValueNode{}
    if err := node.SetRaw(v); err != nil {
        return nil, err
    }

    return node, nil
}

type ValueNode struct {
    data interface{}
}

func (this *ValueNode) GetRaw() (interface{}) {
    return this.data
}

func (this *ValueNode) SetRaw(v interface{}) error {
    switch v.(type) {
    case string, int, int32, int64, float32, float64, bool, nil:
        this.data = v
        return nil
    }

    return fmt.Errorf("Invalid data type for value node")
}

func (this *ValueNode) String() (string, error) {
    if v, ok := this.data.(string); ok {
        return v, nil
    }

    return "", fmt.Errorf("Node value is not string")
}

func (this *ValueNode) MustString() (string) {
    v, err := this.String()
    if err != nil {
        panic(err)
    }

    return v
}

func (this *ValueNode) Int() (int, error) {
    if v, ok := this.data.(int); ok {
        return v, nil
    }

    return 0, fmt.Errorf("Node value is not int")
}

func (this *ValueNode) MustInt() int {
    v, err := this.Int()
    if err != nil {
        panic(err)
    }

    return v
}

func (this *ValueNode) Int32() (int32, error) {
    if v, ok := this.data.(int32); ok {
        return v, nil
    }

    return 0, fmt.Errorf("Node value is not int32")
}

func (this *ValueNode) MustInt32() int32 {
    v, err := this.Int32()
    if err != nil {
        panic(err)
    }

    return v
}

func (this *ValueNode) Int64() (int64, error) {
    return convertToInt64(this.data)
}

func (this *ValueNode) MustInt64() int64 {
    v, err := this.Int64()
    if err != nil {
        panic(err)
    }

    return v
}

func (this *ValueNode) Float32() (float32, error) {
    if v, ok := this.data.(float32); ok {
        return v, nil
    }

    return 0, fmt.Errorf("Node value is not float32")
}

func (this *ValueNode) MustFloat32() (float32) {
    v, err := this.Float32()
    if err != nil {
        panic(err)
    }

    return v
}

func (this *ValueNode) Float64() (float64, error) {
    if v, ok := this.data.(float64); ok {
        return v, nil
    }

    return 0, fmt.Errorf("Node value is not float64")
}

func (this *ValueNode) MustFloat64() (float64) {
    v, err := this.Float64()
    if err != nil {
        panic(err)
    }

    return v
}

func (this *ValueNode) Bool() (bool, error) {
    if v, ok := this.data.(bool); ok {
        return v, nil
    }

    return false, fmt.Errorf("Node value is not bool")
}

func (this *ValueNode) MustBool() (bool) {
    v, err := this.Bool()
    if err != nil {
        panic(err)
    }

    return v
}

func (this *ValueNode) IsNull() bool {
    return this.data == nil
}