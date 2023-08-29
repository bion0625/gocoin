package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/gocoin/explorer"
	"github.com/gocoin/rest"
)

func usage() {
	fmt.Printf("Welcome to Go Coin\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port:	This starts the Port of the server\n")
	fmt.Printf("-mode:	Choose between 'html' and 'rest'\n")
	os.Exit(0)
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	case "server":
		go rest.Start(*port)
		explorer.Start(*port+1)
	default:
		usage()
	}
}
