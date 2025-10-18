package belajargolangjson

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("customer_out.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Ucup",
		MiddleName: "Surucup",
		LastName:   "Tarkucupkucup",
	}

	encoder.Encode(customer)
}
