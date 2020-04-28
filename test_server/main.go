package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	// the scheme was http
	fmt.Println(r.TLS == nil, r.RequestURI)

	_, _ = w.Write([]byte(fmt.Sprintf("TLS, %t url %s", r.TLS == nil, r.URL.String())))
}

func main() {
	http.HandleFunc("/", handler)
	go func() {
		log.Fatal(http.ListenAndServeTLS(":8443", "./localhost.crt", "./localhost.key", nil))
	}()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
