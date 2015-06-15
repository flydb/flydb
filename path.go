package jsondb

import (
    "strings"
)

func parsePath(path string) []string {
    path = strings.Trim(path, ".")
    if path == "" {
        return nil
    }

    return strings.Split(path, ".")
}