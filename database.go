// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package flydb

import (
    "fmt"
    "io/ioutil"
    "path/filepath"
    "time"
)

type Database struct {
    config Config
    root *Node
    isOpen bool
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
        SaveInterval: 5000,
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
    this.isOpen = true
    if this.config.SaveInterval > 0 {
        this.loopSave()
    }

    return nil
}

// Close database
func (this *Database) Close() error {
    this.isOpen = false
    return this.Save()
}

func (this *Database) loopSave() {
    time.AfterFunc(time.Duration(this.config.SaveInterval) * time.Millisecond, func() {
        if this.isOpen {
            this.Save()
            this.loopSave();
        }
    })
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