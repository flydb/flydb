package flydb

type MapNode struct {
    data map[string]*Node
}

func NewMapNode(v map[string]interface{}) (*MapNode, error) {
    node := &MapNode {
    }

    err := node.SetRaw(v); err != nil {
        return nil, err
    }

    return node, nil
}

func (this *MapNode) SetRaw(rawMap map[string]interface{}) (error) {
    data := make(map[string]*Node)
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

func (this *MapNode) GetRaw() map[string]interface{} {
    result := make(map[string]interface{})
    for k, v := range this.data {
        result[k] = v.GetRaw()
    }

    return result
}

func (this *MapNode) Get(k string) (*Node, bool) {
    node, ok := this.data[k]
    return node, ok
}

func (this *MapNode) Set(k string, node *Node) {
    this.data[k] = node
}

func (this *MapNode) Has(k string) bool {
    _, ok := this.data[k]
    return ok
}

func (this *MapNode) Delete(k string) {
    delete(this.data[k])
}