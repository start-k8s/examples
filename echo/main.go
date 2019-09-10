package main

import (
        "fmt"
        "log"
        "net/http"
)

func main() {
        http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
                log.Println("received request")
                fmt.Fprintf(writer, "Hello, Docker!!")
        })

        log.Println("start server")

        server := &http.Server{
                Addr: ":5000",
        }
        if err := server.ListenAndServe(); err != nil {
                log.Println(err)
        }
}
