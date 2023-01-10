package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	var port = flag.String("port", "8080", "")
	flag.Parse()

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	fs := http.FileServer(http.Dir(cwd))
	http.Handle("/", fs)

	fmt.Println("server start with " + *port)
	http.ListenAndServe(":"+*port, nil)
}
