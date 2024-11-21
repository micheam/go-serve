package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	flag.Usage = usage
	flag.Parse()
	ctx := context.Background()

	if flag.NArg() == 1 {
		dir = flag.Arg(0)
	}
	err := run(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}
}

func usage() {
	usage := `Usage: serve [options] [dir]
Serve the current directory or the specified directory over HTTP.

Options:
  -port string port to listen on (default "8080")

Examples:
  serve current directory on port 8080
  $ serve

  serve a specific directory on port 8081
  $ serve -port 8081 /path/to/directory
`
	fmt.Fprintln(os.Stderr, usage)
	os.Exit(2)
}

var (
	port = flag.String("port", "8080", "")
	dir  = "."
)

func run(_ context.Context) error {
	if dir != "." {
		err := os.Chdir(dir)
		if err != nil {
			return fmt.Errorf("failed to change directory: %v", err)
		}
	}
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %v", err)
	}

	fs := http.FileServer(http.Dir(cwd))
	http.Handle("/", fs)

	fmt.Println("serving " + cwd)
	fmt.Println("listening on port " + *port)

	return http.ListenAndServe(":"+*port, nil)
}
