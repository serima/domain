package main

import (
	"fmt"
	"os"

	"github.com/domainr/whois"
)

func main() {
	query := "domai.nr"
	request, err := whois.NewRequest(query)
	FatalIf(err)
	response, err := whois.DefaultClient.Fetch(request)
	FatalIf(err)
	body := response.String()
	fmt.Println(body)
}

func FatalIf(err error) {
	if err == nil {
		return
	}
	fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	os.Exit(-1)
}
