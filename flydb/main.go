package main

import (
    "os"
    "github.com/codegangsta/cli"
    "github.com/flydb/flydb"
    "regexp"
)

func checkListen(listen string) bool {
    var simpleIp = "(\\d+\\.){3}\\d+"
    var simplePort = ":\\d+"
    for _, v := range []string {
        simpleIp,
        simplePort,
        simpleIp + simplePort,
    } {
        if ok, _ := regexp.MatchString("^" + v + "$", listen); ok {
            return true
        }
    }

    return false
}

func main() {
    app := cli.NewApp()
    app.Name = "flydb"
    app.Usage = "FlyDB HTTP server"
    app.Action = func(c *cli.Context) {
        var file string
        var listen = "127.0.0.1:8080"

        argLength := len(c.Args())
        if (argLength == 0) {
        } else if argLength == 1 {
            arg := c.Args()[0]

            if checkListen(arg) {
                listen = arg
            } else {
                file = arg
            }
        } else {
            file = c.Args()[0]
            listen = c.Args()[1]
        }

        var db *flydb.Database
        var err error
        if file == "" {
            db = flydb.Memory()
        } else {
            db , err = flydb.Open(file)
            if err != nil {
                panic(err)
            }
        }

        server := &flydb.Server {
            db,
        }

        println("Start server: " + listen)
        if err := server.Serve(listen); err != nil {
            panic(err)
        }
    }

    app.Run(os.Args)
}