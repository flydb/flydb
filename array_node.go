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

func (this *ArrayNode) Set(i int, n *Node) {

}

func (this *ArrayNode) Delete(key int) {
    if key < 0 || key >= len(this.data) {
        return
    }

    this.data = append(this.data[0:key], this.data[key+1:])
}

func (this *ArrayNode) Length() int {
    return len(this.data)
}