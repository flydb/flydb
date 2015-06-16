package jsondb

import (
    "fmt"
)

type ValueNode struct {
    data interface{}
}

func (this *ValueNode) GetRaw() (interface{}) {
    return this.data
}

func (this *ValueNode) SetRaw(v interface{}) error {
    switch typedValue = v.(type) {
    case string, int, float32, bool, nil:
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
    if err {
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

func (this *ValueNode) MustInt() (int, error) {
    v, err := this.Int()
    if err {
        panic(err)
    }

    return v
}

func (this *ValueNode) Float() (float32, error) {
    if v, ok := this.data.(float32); ok {
        return v, nil
    }

    return 0, fmt.Errorf("Node value is not float")
}

func (this *ValueNode) MustFloat() (float32) {
    v, err := this.Float()
    if err {
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

func (this *ValueNode) MustBool() (bool, error) {
    v, err := this.Bool()
    if err {
        panic(err)
    }

    return v
}

func (this *ValueNode) IsNull() bool {
    
}