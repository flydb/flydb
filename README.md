# flydb

[![Travis](https://img.shields.io/travis/flydb/flydb.svg?style=flat-square)](https://travis-ci.org/flydb/flydb)
[![Go doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/flydb/flydb)
[![GitHub license](https://img.shields.io/github/license/flydb/flydb.svg?style=flat-square)](LICENSE)

    In progress...

Pure Go database with simple format.

## Features

- Support different data file format: `JSON`, `YAML`, `XML`
- Powerful access API

## Limits

- Performance: flydb is not built for performance, so you may not want to use it with huge data

## Usage

### Embed in Golang

```go
package main

import (
    "github.com/flydb/flydb"
    "log"
)

func main() {
    db, err := flydb.Open("/path/to/db.json")
    if err != nil {
        fmt.Errorf("cannot open database")
    }
    email := db.Root().MustGet("users.3.email").MustVallue().MustString()
    log.Println(email)
}
```

### HTTP Server

See `examples/server/main.go`.

### HTTP API

#### Get

```
curl http://127.0.0.1:8080/users/1/email
```

#### Set

```
curl -x PUT -D '{"key": "value"}' http://127.0.0.1:8080/users/1/metadata
```

#### Delete

```
curl -x DELETE http://127.0.0.1:8080/users/1/metadata
```

## Internal

When you are using `flydb` as a Golang library, it's very helpful to understand it's internal.

### Node

Data is saved as a tree like structure in memory, you can find and modify every node in the tree.

We only care about three types of data:

- map: `map[string]interface{}` in Golang
- array: `[]interface{}` in Golang
- value: `string`, `int`, `float32`, `bool`, `nil` in Golang

#### Working with Node

Create a node from raw data:

```go
rootNode, err := flydb.CreateNode(rawdata)
```

To get the value of specific node, you can:

1. Find the node

    ```
    node, err := rootNode.Get("books.3.title")
    ```

2. Assert node type

    ```go
    valueNode, err := node.Value()
    // arrayNode, err := node.Array()
    // mapNode, err := node.Map()
    ```

3. Assert exact value type

    ```go
    title, err := typedNode.String()
    ```

If you know exact structure of your data, you can simply write in one line:

```go
title := rootNode.MustGet("books.3.title").MustValue().MustString()
```

To update value of specific node, you can use the `Set` method:

```go
rootNode.Set("books.3.date", "2015-01-01")
```

### Format

The in memory node data can be transfered between different formats: JSON, YAML, XML and many other custom formats.

Load from formated data:

```go
jsonData := `
{
    "key": "value",
    "key1": "value1"
}
`

format := flydb.GetFormat("json")

rawData, err := format.Unmarshal(jsonData)
node, err := flydb.CreateNode(rawData)
```

Export node with format:

```go
rawData := node.GetRaw()
format := flydb.GetFormat("yaml")
yamlData, err := format.Marshal(rawData)
println(string(yamlData))
```

## TODO

- XML format support
- Query API