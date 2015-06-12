package jsondb

type Node interface {
    Marshal() []byte
    Unmarshal(b []byte) error
}

type Node {
    data []byte
}

// Get child node by path
func (this *Node) Get(path string) (*Node, bool) {
    return nil, false
}

func (this *Node) MustGet(path string) *Node {
    panic("TODO")
    return nil
}

// Check existence of child node
func (this *Node) Has(key string) *Node bool {

}

// Set value of child node by path
func (this *Node) Set(key string, v interface{}) error {

}

func (this *Node) IsArray() bool {
    
}

func (this *Node) IsMap() bool {

}

func (this *Node) String() string {

}

func (this *Node) Int() int {

}

func (this *Node) Array() ([]interface{}, error) {

}

func (this *Node) Map() (map[string]interface{}, error) {

}

func (this *Node) Unmarshal(v) error {

}

func (this *Node) Marshal() string {

}

type ArrayNode struct {
    data []*Node
}

func (this *ArrayNode) Marshal() {

}

func (this *ArrayNode) Unmarshal() {

}

type MapNode struct {
    data map[string]*Node
}

type ValueNode struct {
    data interface{}
}
