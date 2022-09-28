package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var addr = flag.String("http-port", ":8080", "The address to listen on for HTTP requests.")

func main() {
	flag.Parse()

	aws_secret := "AKIAIMNOJVGFDXXXE4OA"
fmt.Println(aws_secret)
	fmt.Println("starting prometheus metric server")
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}
