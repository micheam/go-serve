# serve [![Go](https://github.com/micheam/serve/actions/workflows/go.yml/badge.svg)](https://github.com/micheam/serve/actions/workflows/go.yml)
Tiny Program, which serves the current dir on HTTP GET.

## Usage
```
Usage: serve [options] [dir]
Serve the current directory or the specified directory over HTTP.

Options:
  -port string port to listen on (default "8080")

Examples:
  serve current directory on port 8080
  $ serve

  serve a specific directory on port 8081
  $ serve -port 8081 /path/to/directory
```

## Installation
Install via the `go install` command:

    $ go install github.com/micheam/serve@latest
    $ serve -port 6000 /path/to/directory

or, Run directly via the `go run`:

    $ go run github.com/micheam/serve@latest -port 6000 .

## License
MIT

## Author
[Michito Maeda](https://github.com/micheam)
