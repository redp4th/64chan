package main

import (
	"net/http"
	"nmb/router"
)

func main() {
	r := router.InitRouter()
	s := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}
	// log.Println(s.ListenAndServeTLS("cert/cert.pem", "cert/key.pem"))
	s.ListenAndServe()
}
