package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/likexian/whois-go"
	"github.com/likexian/whois-parser-go"
)

type Response struct {
	DomainStatus   string `json:"domainStatus"`
	CreatedDate    string `json:"createdDate"`
	ExpirationDate string `json:"expirationDate"`
	Name           string `json:"name"`
	Email          string `json:"email"`
}

func main() {
	whoisRaw, err := whois.Whois("serima.co")
	FatalIf(err)
	result, err := whois_parser.Parser(whoisRaw)
	if err == nil {
		res2D := &Response{
			DomainStatus:   result.Registrar.DomainStatus,
			CreatedDate:    result.Registrar.CreatedDate,
			ExpirationDate: result.Registrar.ExpirationDate,
			Name:           result.Registrant.Name,
			Email:          result.Registrant.Email}
		res2B, _ := json.Marshal(res2D)
		fmt.Println(string(res2B))
	}
}

func FatalIf(err error) {
	if err == nil {
		return
	}
	fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	os.Exit(-1)
}
