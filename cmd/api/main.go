package main

import (
	"net/http"
	"github.com/tantanok221/douren-backend/handler"
)

func main() {
	mux := http.NewServeMux()
	print("Server Running")
	mux.Handle("/artist", handler.GetAllArtist())
	http.ListenAndServe(":3000",mux)
}
