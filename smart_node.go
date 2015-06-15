package jsondb

type SmartNode struct {
    node *Node
}

func (this *SmartNode) Marshal() ([]byte, error) {
    return this.node.Marshal()
}

func (this *SmartNode) Unmarshal(b []byte) error {
    return this.node.Unmarshal(b)
}

func (this *SmartNode) Get(path) (*SmartNode, bool) {
    
}

func (this *SmartNode) Set(path, node *SmartNode) (error) {

}

func (this *SmartNode) Delete(path) error {

}