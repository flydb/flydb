// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package flydb

import (
    "net/http"
    "fmt"
)

func Serve(path, listen string) error {
    db, err := Open(path)
    if err != nil {
        return err
    }

    server := &Server{
        db,
    }
    return server.Serve(listen)
}

type Server struct {
    database *Database
}

func (this *Server) Serve(listen string) error {
    http.HandleFunc("/get", this.HandleGet)
    http.HandleFunc("/set", this.HandleSet)
    return http.ListenAndServe(listen, nil)
}

func (this *Server) HandleGet(w http.ResponseWriter, r *http.Request) {
    path := r.FormValue("path")
    node, err := this.database.Root().Get(path)
    if err != nil {
        panic(err)
    }

    format := GetFormat("json")
    bytes, err := format.Marshal(node.GetRaw())
    if err != nil {
        panic(err)
    }

    fmt.Fprint(w, string(bytes))
}

func (this *Server) HandleSet(w http.ResponseWriter, r *http.Request) {

}