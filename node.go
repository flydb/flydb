package jsondb

type Node {

}

// Get child node by path
func (this *Node) Get(path string) *Node {

}

// Check existence of child node
func (this *Node) Has(key string) *Node bool {

}

// Set value of child node by path
func (this *Node) Set(key string, v interface{}) error {

}

func (this *Node) ToJSON() []byte {

}

func (this *Node) LoadJSON(v []byte) {

}

func (this *Node) String() string {

}

func (this *Node) Int() int {

}