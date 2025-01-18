package main

import (
	"encoding/json"
	"fmt"
)

type Customer struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	customer := Customer{
		Name: "Lorem",
		Age:  26,
	}
	
	// Serialization (Go struct to JSON string)
	jsonBytes, err := json.Marshal(customer)

	if err != nil {
		fmt.Println(err)
	}

	customerSerialized := string(jsonBytes)

	fmt.Println(string(customerSerialized))

	// Deserialization (JSON string to Go struct)
	var customerDeserialized Customer

	jsonString := `{"Name":"Lorem","Age":26}`
	err = json.Unmarshal([]byte(jsonString), &customerDeserialized)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(customerDeserialized)
}

