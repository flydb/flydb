// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package flydb

import (
    "net/http"
    "io/ioutil"
    "fmt"
    "strings"
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
    Database *Database
}

func (this *Server) Serve(listen string) error {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "GET":
            this.DoGet(w, r)
            break
        case "PUT":
            this.DoSet(w, r)
            break
        case "DELETE":
            this.DoDelete(w, r)
        default:
            http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
        }
    })

    return http.ListenAndServe(listen, nil)
}

func (this *Server) DoGet(w http.ResponseWriter, r *http.Request) {
    path := this.GetRequestPath(r)

    node, err := this.Database.Root().Get(path)
    if err != nil {
        this.Error(w, err)
        return
    }

    format := GetFormat("json")
    println(node.GetRaw())
    bytes, err := format.Marshal(node.GetRaw())
    if err != nil {
        this.Error(w, err)
        return
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    fmt.Fprint(w, string(bytes))
}

func (this *Server) DoSet(w http.ResponseWriter, r *http.Request) {
    path := this.GetRequestPath(r)

    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        panic(err)
    }

    jsonFormat := GetFormat("json")

    v, err := jsonFormat.Unmarshal(body)
    if err != nil {
        this.Error(w, err)
        return
    }

    if err := this.Database.Root().Set(path, v); err != nil {
        this.Error(w, err)
        return
    }
}

func (this *Server) DoDelete(w http.ResponseWriter, r *http.Request) {
    path := this.GetRequestPath(r)

    if err := this.Database.Root().Delete(path); err != nil {
        this.Error(w, err)
        return
    }
}

func (this *Server) Error(w http.ResponseWriter, err error) {
    http.Error(w, err.Error(), 500)
}

func (this *Server) GetRequestPath(r *http.Request) string {
    return strings.Replace(r.URL.Path, "/", ".", -1)
}