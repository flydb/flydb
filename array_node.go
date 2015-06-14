package jsondb

type ArrayNode struct {
    data []*Node
}

func (this *ArrayNode) Marshal() ([]byte, error) {

}

func (this *ArrayNode) Unmarshal(b []byte) error {

}

func (this *ArrayNode) Append(n *Node) {

}

func (this *ArrayNode) Get(i int) {
    
}