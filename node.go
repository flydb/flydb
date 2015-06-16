package jsondb

type Node interface {
    // Marshal() []byte
    // Unmarshal(b []byte) error
    SetRaw(interface{}) error
    GetRaw() interface{}
}