package flydb

import (
    "fmt"
    "io/ioutil"
    "path/filepath"
)

type Database struct {
    path string
    format Format
    root *SmartNode
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
    var realFormat Format
    switch typedFormat := format.(type) {
    case string:
        realFormat = GetFormat(typedFormat)
        if realFormat == nil {
            fmt.Errorf("unknown format: %s", typedFormat)
        }
    case Format:
        realFormat = typedFormat
    default:
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

    node, err := CreateNodeFromRawData(v)
    if err != nil {
        return nil, err
    }

    root := &SmartNode {
        node,
    }

    return &Database{
        path: path,
        root: root,
    }, nil
}

// Create an in memory database
func Memory() *Database {
    return &Database {
        root: &SmartNode {
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
    raw := this.root.GetRaw()
    b, err := this.format.Marshal(raw)
    if err != nil {
        return err
    }

    return ioutil.WriteFile(path, b, 0600)
}

// Check key existence
func (this *Database) Has(path interface{}) bool {
    return this.root.Has(path)
} 

// Get node by key
func (this *Database) Get(path interface{}) (*SmartNode, error) {
    return this.root.Get(path)
}

func (this *Database) MustGet(path interface{}) (*SmartNode) {
    return this.root.MustGet(path)
}

// Set node value
func (this *Database) Set(path interface{}, v interface{}) error {
    return this.root.Set(path, v)
}