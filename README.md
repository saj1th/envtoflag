# envtoflag [![GoDoc](https://godoc.org/github.com/saj1th/envtoflag?status.png)](https://godoc.org/github.com/saj1th/envtoflag)

envtoflag is a lib to populate flags from environment variables (to support http://12factor.net/config)

## USage

```go
package main

import (
	"flag"
	"log"

	"github.com/saj1th/envtoflag"
)

func main() {
	var (
		mode string // dev|debug|prod
		port int    // port to listen to
	)

	flag.StringVar(&mode, "mode", "", "dev|debug|prod")
	flag.IntVar(&port, "port", 8080, "port to listen to")
	
	// export YOURPACKAGENAME_PORT=9090	
	envtoflag.Parse("yourpackagename")

	fmt.Printf("Mode: %s\nPort: %d\n\n", mode, port)

}

```