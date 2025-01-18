package main

import (
	"encoding/json"
	"fmt"
	"structs/creating_structs"
	"structs/embedded_structs"
	"structs/struct_tags"
	"time"
)

func main() {
	customer1 := creating_structs.NewCustomer("Lorem", "Ipsum", time.Now().Format("02/01/2006"))

	customer1.OutputUserDetails()

	// embedded structs
	admin := embedded_structs.Admin{
		Email:    "admin@example.com",
		Password: "password123",
		User: embedded_structs.User{
			FirstName: "John",
			LastName:  "Doe",
		},
	}

	fmt.Println(admin.Greet()) // Output: Hello, John Doe!

	// struct tags
	note := struct_tags.Note{
		Title: "my title",
		Content: "my content",
		CreatedAt: time.Now(),
	}

	// Convert struct to JSON
	jsonData, err := json.Marshal(note)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	// Print the JSON string
	fmt.Println(string(jsonData))
}