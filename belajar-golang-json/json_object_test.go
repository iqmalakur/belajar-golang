package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hobbies    []string
	Addresses  []Address
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Ucup",
		MiddleName: "Bin",
		LastName:   "Surucup",
		Age:        10,
		Married:    true,
	}

	bytes, _ := json.Marshal(customer)

	fmt.Println(string(bytes))

	bytes, _ = json.MarshalIndent(customer, "", "  ")

	fmt.Println(string(bytes))
}
