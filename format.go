// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

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

func GuessFormat(format interface{}) Format {
    var realFormat Format
    switch typedFormat := format.(type) {
    case string:
        realFormat = GetFormat(typedFormat)
    case Format:
        realFormat = typedFormat
    default:
        return nil
    }

    return realFormat
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