package jsondb

type MapNode struct {
    data map[string]*Node
}

func (this *MapNode) Marshal() ([]byte, error) {

}

func (this *MapNode) Unmarshal(b []byte) error {

}

func (this *MapNode) Get(k string) (*Node, bool) {

}

func (this *MapNode) Set(k string, n *Node) {

}

func (this *MapNode) Delete(k string) {
    
}