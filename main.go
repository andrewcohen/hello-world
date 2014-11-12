package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

var port = flag.String("p", "3000", "port number")

func main() {
	flag.Parse()

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hello %s", strings.Join(os.Args[1:], ","))
	})

	log.Printf("Listening on port %s", *port)
	http.ListenAndServe(":"+*port, nil)
}
