package flydb

import (
    "strings"
)

func parsePath(path interface{}) []string {
    if path == nil {
        return nil
    }
    switch typedPath := path.(type) {
    case []string:
        return typedPath
    case string:
        typedPath = strings.Trim(typedPath, ".")
        if typedPath == "" {
            return nil
        }

        return strings.Split(typedPath, ".")
    default:
        panic("invalid path")
    }
}