package main

import (
	"flag"
	"log"

	"github.com/s02190058/reverse-proxy-cache-client/internal/app"
)

var (
	host  = flag.String("host", "localhost", "service host")
	port  = flag.String("port", "9090", "service port")
	dir   = flag.String("dir", ".", "download directory")
	async = flag.Bool("async", false, "make asynchronous calls")
)

// main is a client entrypoint.
func main() {
	flag.Parse()

	a, err := app.New(*host, *port, *dir, *async)
	if err != nil {
		log.Fatal(err)
	}

	a.Run()
}
