package jsondb

type Database struct {
    path string
    root *SmartNode
}

func Open(path) *Database {
    return &Database{
        path: path,
    }
}

// Create an in memory database
func Memory() *Database {
    return &Database {

    }
}

// Close database
func (this *Database) Close() {

}

// Flush changes to disk
func (this *Database) Flush() {

}

// Save database as another file
func (this *Database) SaveAs(path string) {

}

// Convert database to JSON
func (this *Database) ToJSON() []byte {

}

// Load database from JSON
func (this *Database) LoadJSON([]byte) error {

}

// Load a JSON file
func (this *Database) LoadFile(path) error {

}

// Check key existence
func (this *Database) Has(path string) bool {

} 

// Get node by key
func (this *Database) Get(path string) (*SmartNode, bool) {
    return this.root.Get(path)
}

// Set node value
func (this *Database) Set(path string, v interface{}) {
    return this.root.Set(path, v)
}