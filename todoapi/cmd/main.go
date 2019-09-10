package main

import (
	"fmt"
	"github.com/gihyodocker/todoapi"
	"log"
	"net/http"
	"os"
)

func main() {
	env, err := todoapi.CreateEnv()
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}

	masterDB, err := todoapi.CreateDbMap(env.MasterURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s is invalid database", env.MasterURL)
		return
	}

	slaveDB, err := todoapi.CreateDbMap(env.SlaveURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s is invalid database", env.SlaveURL)
		return
	}

	mux := http.NewServeMux()

	todoHandler := todoapi.NewTodoHandler(masterDB, slaveDB)
	mux.Handle("/todo", todoHandler)
	mux.HandleFunc("/hc", hc)


	server := http.Server{
		Addr:    env.Bind,
		Handler: mux,
	}

	log.Printf("Listen HTTP Server")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func hc(writer http.ResponseWriter, request *http.Request) {
	log.Println("[GET] /hc")
	writer.Write([]byte("OK"))
}
