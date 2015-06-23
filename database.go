// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package flydb

import (
    "fmt"
    "io/ioutil"
    "path/filepath"
)

type Database struct {
    config Config
    root *Node
}

func New(config Config) *Database {
    config = setDefaultConfig(config)
    return &Database{
        config: config,
    }
}

func Open(path string) (*Database, error) {
    db := New(Config {
        Path: path,
        Save: true,
        SaveInterval: 5,
    })
    if err := db.Open(); err != nil {
        return nil, err
    }

    return db, nil
}

// Create an in memory database
func Memory() *Database {
    return &Database {
        root: &Node {
        },
    }
}

func (this *Database) Open() error {
    if this.config.Format == nil {
        return fmt.Errorf("unknown format")
    }

    bytes, err := ioutil.ReadFile(this.config.Path)
    if err != nil {
        return err
    }

    v, err := this.config.Format.Unmarshal(bytes)
    if err != nil {
        return err
    }

    root, err := CreateNode(v)
    if err != nil {
        return err
    }

    this.root = root

    return nil
}

// Close database
func (this *Database) Close() error {
    return this.Save()
}

// Flush changes to disk
func (this *Database) Flush() {

}

func (this *Database) Save() error {
    return this.SaveAs(this.config.Path)
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