package main

import (
	"fmt"
)

func main() {
	fmt.Println(`Usage of gurl:
	-d string
		  HTTP request body
	-h    Show this help message
	-i    Include HTTP response headers
	-m string
		  HTTP method (default "GET")
	-o string
		  Write HTTP response body to file
	-v    Verbose output
	-x string
		  HTTP proxy
	-X string
		  HTTP method (default "GET")`)
}
