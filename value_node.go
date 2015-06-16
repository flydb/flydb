package jsondb

type ValueNode struct {
    data interface{}
}

func (this *ValueNode) Marshal() ([]byte, error) {

}

func (this *ValueNode) Unmarshal([]byte) error {

}

func (this *ValueNode) SetValue(value interface{}) error {
    this.data = value
    return nil
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