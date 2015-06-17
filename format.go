package flydb

type Format interface {
    Marshal(v interface{}) ([]byte, error)
    Unmarshal(b []byte) (v interface{}, error)
}