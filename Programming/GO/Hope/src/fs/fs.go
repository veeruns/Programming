package main

import (
	"flag"
	"net/http"
	"os"
)

func main() {
	var dir string
	path := flag.String("path", "", "Path of directory to serve")
	port := flag.String("port", "3000", "Port to Server HTTP File server On")
	flag.Parse()
	if *path == "" {
		dir, _ = os.Getwd()
	} else {
		dir = *path
	}

	http.ListenAndServe(":"+*port, http.FileServer(http.Dir(dir)))
}
