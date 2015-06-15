package jsondb

type ValueNode struct {

}

func (this *ValueNode) Marshal() ([]byte, error) {

}

func (this *ValueNode) Unmarshal([]byte) error {

}

func (this *ValueNode) String() (string, error) {

}

func (this *ValueNode) MustString() (string) {

}

func (this *ValueNode) Int() (int, error) {

}

func (this *ValueNode) MustInt() (int, error) {

}

func (this *ValueNode) Float() (float32, error) {

}

func (this *ValueNode) MustFloat() (float32) {

}

func (this *ValueNode) Bool() (bool, error) {

}

func (this *ValueNode) MustBool() (bool, error) {

}

func (this *ValueNode) IsNull() bool {

}