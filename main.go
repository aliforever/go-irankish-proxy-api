package main

import (
	"net/http"

	"github.com/aliforever/go-irankish-proxy-api/handlers"
)

func main() {
	http.HandleFunc("/ik_make_token", handlers.MakeTokenRequestHandler)
	http.HandleFunc("/ik_verify_payment", handlers.VerifyPaymentRequestHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
