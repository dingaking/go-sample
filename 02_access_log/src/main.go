/*
https://github.com/mash/go-accesslog
*/

package main

import (
    "log"
    "net/http"

    accesslog "github.com/mash/go-accesslog"
)

type logger struct {
}

func (l logger) Log(record accesslog.LogRecord) {
    log.Println(record.Method + " " + record.Uri)
}

func main() {
    l := logger{}
    handler := http.FileServer(http.Dir("."))
    http.ListenAndServe(":8888", accesslog.NewLoggingHandler(handler, l))
}
