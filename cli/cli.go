package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/dalyoon-practice/nomadcoin/explorer"
	"github.com/dalyoon-practice/nomadcoin/rest"
)

func printUsage() {
	fmt.Printf("Welcome to Nomad Coin\n\n")
	fmt.Printf("Please use the following commands:\n\n")
	fmt.Printf("-port 4000:	Set the PORT of server\n")
	fmt.Printf("-mode html:		Choose between 'html' and 'rest'\n\n")
	runtime.Goexit()
}

func Start() {
	if len(os.Args) < 2 {
		printUsage()
	}

	port := flag.Int("port", 4000, "Set port of server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()
	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		printUsage()
	}
}
