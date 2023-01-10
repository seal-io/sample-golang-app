package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var addr = flag.String("http-port", ":8080", "The address to listen on for HTTP requests.")

func main() {
	flag.Parse()

	fmt.Println("starting prometheus metric server")
	log.Fatal(http.ListenAndServe(*addr, nil))
}
