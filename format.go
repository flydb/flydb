package flydb

import (
    "strings"
)

type Format interface {
    Extensions() []string
    Marshal(v interface{}) ([]byte, error)
    Unmarshal(b []byte) (interface{}, error)
}

var formatMap map[string]Format

func RegisterFormat(name string, format Format) {
    formatMap[name] = format
}

func GetFormat(name string) Format {
    return formatMap[name]
}

func CheckFormatByExtension(ext string) Format {
    ext = strings.ToLower(strings.TrimLeft(ext, "."))
    for _, format := range formatMap {
        for _, v := range format.Extensions() {
            if ext == strings.ToLower(strings.TrimLeft(v, ".")) {
                return format
            }
        }
    }
    
    return nil
}

func init() {
    formatMap = make(map[string]Format)
}