package main

import (
	"fmt"
	"net/http"
)

const (
	portNumber = "8888"
)

func main() {

	fmt.Printf("App starting on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)

}
