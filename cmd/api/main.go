package main

import (
    "net/http"
    "fmt"
    "github.com/go-chi/chi"
    "github.com/kaus19/online_offline_tracker/internal/handlers"

    log "github.com/sirupsen/logrus"

)

func main() {
    log.SetReportCaller(true)
    var r *chi.Mux = chi.NewRouter()

    handlers.Handler(r)

    fmt.Println("Starting Go-API Service!!")

    err := http.ListenAndServe("localhost:8000",r)
    if err!=nil{
        log.Error(err)
    }
}
