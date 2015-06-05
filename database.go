package jsondb

type Database struct {
    path string
    root *Node
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
func (this *Database) Has(key string) bool {

} 

// Get node by key
func (this *Database) Get(key string) *Node {

}

// Set node value
func (this *Database) Set(key string, v interface{}) {

}