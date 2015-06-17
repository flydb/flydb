package flydb

type Format interface {
    CheckFormat(name string) bool
    Marshal(v interface{}) ([]byte, error)
    Unmarshal(b []byte) (v interface{}, error)
}