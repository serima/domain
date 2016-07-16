package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/likexian/whois-go"
	"github.com/likexian/whois-parser-go"
)

// Response : struct of json response
type Response struct {
	DomainStatus   string `json:"domainStatus"`
	CreatedDate    string `json:"createdDate"`
	ExpirationDate string `json:"expirationDate"`
	Name           string `json:"name"`
	Email          string `json:"email"`
}

func main() {
	var domain = flag.String("d", "google.com", "domain")
	flag.Parse()

	whoisRaw, err := whois.Whois(*domain)
	FatalIf(err)
	result, err := whois_parser.Parser(whoisRaw)
	FatalIf(err)
	res := &Response{
		DomainStatus:   result.Registrar.DomainStatus,
		CreatedDate:    result.Registrar.CreatedDate,
		ExpirationDate: result.Registrar.ExpirationDate,
		Name:           result.Registrant.Name,
		Email:          result.Registrant.Email}
	resJSON, _ := json.Marshal(res)
	fmt.Println(string(resJSON))
}

// FatalIf : exit, if error occurs
func FatalIf(err error) {
	if err == nil {
		return
	}
	fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	os.Exit(-1)
}
