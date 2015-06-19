// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package flydb

import (
    "fmt"
    "io/ioutil"
    "path/filepath"
)

type Database struct {
    path string
    format Format
    root *Node
}

func Open(path string) (*Database, error) {
    ext := filepath.Ext(path)

    format := CheckFormatByExtension(ext)
    if format == nil {
        return nil, fmt.Errorf("unknown format")
    }

    return OpenFormat(path, format)
}

func OpenFormat(path string, format interface{}) (*Database, error) {
    var realFormat = GuessFormat(format)
    if realFormat == nil {
        return nil, fmt.Errorf("unknown format")
    }

    bytes, err := ioutil.ReadFile(path)
    if err != nil {
        return nil, err
    }

    v, err := realFormat.Unmarshal(bytes)
    if err != nil {
        return nil, err
    }

    root, err := CreateNode(v)
    if err != nil {
        return nil, err
    }

    return &Database{
        path: path,
        root: root,
    }, nil
}

// Create an in memory database
func Memory() *Database {
    return &Database {
        root: &Node {
        },
    }
}

// Close database
func (this *Database) Close() {

}

// Flush changes to disk
func (this *Database) Flush() {

}

// Save database as another file
func (this *Database) SaveAs(path string) error {
    ext := filepath.Ext(path)

    format := CheckFormatByExtension(ext)
    return this.SaveAsFormat(path, format)
}

func (this *Database) SaveAsFormat(path string, format interface{}) error {
    var realFormat = GuessFormat(format)
    if realFormat == nil {
        return fmt.Errorf("unknown format")
    }

    raw := this.root.GetRaw()
    b, err := realFormat.Marshal(raw)
    if err != nil {
        return err
    }

    return ioutil.WriteFile(path, b, 0600)
}

func (this *Database) Root() *Node {
    return this.root
}