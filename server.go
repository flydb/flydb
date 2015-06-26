// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package flydb

import (
    "net/http"
)

func Serve(path, listen string) {

}

type Server struct {
    database *Database
}

func (this *Server) Serve(listen string) {
    http.HandleFunc("/get", this.HandleGet)
    http.HandleFunc("/set", this.HandleSet)
    http.ListenAndServe(listen, nil)
}

func (this *Server) HandleGet(w http.ResponseWriter, r *http.Request) {

}

func (this *Server) HandleSet(w http.ResponseWriter, r *http.Request) {

}