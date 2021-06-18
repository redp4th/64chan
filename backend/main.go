package main

import (
	"log"
	"net/http"
	"nmb/router"
)

func main() {
	r := router.InitRouter()
	s := &http.Server{
		Addr:           ":8000",
		Handler:        r,
		MaxHeaderBytes: 16 << 20,
	}
	log.Println(s.ListenAndServeTLS("localhost.crt", "localhost.key"))
	// s.ListenAndServe()
}
