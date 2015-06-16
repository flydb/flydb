package jsondb

type MapNode struct {
    data map[string]*Node
}

func NewMapNode() {
    return &MapNode {
        make(map[string]*Node),
    }
}

func (this *MapNode) Marshal() ([]byte, error) {

}

func (this *MapNode) Unmarshal(b []byte) error {

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