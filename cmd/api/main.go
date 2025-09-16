package main

import (
	"fmt"
	"log"
	"net/http"
	
	"github.com/joho/godotenv"
)

const webPort = "8080"

func main() {
	godotenv.Load()
	
	serve()
}

func serve() {
	log.Println("Server listening on port ", webPort)

	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: NewRouter(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic("Server initialization failed: ", err)
	}
}
