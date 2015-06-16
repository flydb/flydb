package jsondb

import (
    "io/ioutil"
)

type Database struct {
    path string
    root *SmartNode
}

func Open(path string) (*Database, error) {
    bytes, err := ioutil.ReadFile(path)
    if err != nil {
        return nil, err
    }

    root := &SmartNode {
    }
    if err := root.Unmarshal(bytes); err != nil {
        return nil, err
    }

    return &Database{
        path: path,
        root: root,
    }
}

// Create an in memory database
func Memory() *Database {
    return &Database {
        root: &SmartNode {
            NewMapNode(),
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
    b, err := this.Marshal()
    if err != nil {
        return err
    }

    return ioutil.WriteFile(path, b, 0600)
}

// Convert database to JSON bytes
func (this *Database) Marshal() ([]byte, error) {
    return this.root.Marshal()
}

// Load database from JSON
func (this *Database) Unmarshal(b []byte) error {
    return this.root.Unmarshal(b)
}

// Check key existence
func (this *Database) Has(path interface{}) bool {
    return this.root.Has(path)
} 

// Get node by key
func (this *Database) Get(path interface{}) (*SmartNode, bool) {
    return this.root.Get(path)
}

// Set node value
func (this *Database) Set(path interface{}, v interface{}) {
    return this.root.Set(path, v)
}