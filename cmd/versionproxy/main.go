package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/csrwng/versionproxy/pkg/versionproxy"
)

const (
	defaultPort = ":2374"
)

func main() {
	flag.Parse()
	args := os.Args
	var listenSpec string
	if len(args) <= 1 {
		listenSpec = defaultPort
	} else {
		listenSpec = args[1]
	}
	fmt.Printf("Starting version proxy on %s\n", listenSpec)
	err := http.ListenAndServe(listenSpec, versionproxy.New())
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

const usageString = `
Usage:
%[1]s LISTEN_SPEC

where LISTEN_SPEC is either a port (ie. :1080) 
or an IP and port (ie. 127.0.0.1:1080) 

Example:
%[1]s ":2375"
`

func usage() string {
	return fmt.Sprintf(usageString, os.Args[0])
}
