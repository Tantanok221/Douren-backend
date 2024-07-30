package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/tantanok221/douren-backend/internal/routes"
)

func main() {

	errorLog := log.New(os.Stderr, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile,)
	addr := flag.String("addr", ":3000", "HTTP Network address")
	print("Server Running")

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  routes.Route(),
	}
	srv.ListenAndServe()
}
