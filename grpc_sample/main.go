package main

import (
	"fmt"
	"log"
	"vg/groc_sample/examplepb"

	// Import the package from the generated code

	"google.golang.org/protobuf/proto"
)

func main() {
	// Create a new Person message
	person := &examplepb.Person{
		Name:  "John Doe",
		Id:    1234,
		Email: "johndoe@example.com",
	}

	// Serialize the Person message to a binary format
	data, err := proto.Marshal(person)
	if err != nil {
		log.Fatalf("Failed to marshal person: %v", err)
	}

	// Deserialize the binary data back into a Person message
	newPerson := &examplepb.Person{}
	err = proto.Unmarshal(data, newPerson)
	if err != nil {
		log.Fatalf("Failed to unmarshal data: %v", err)
	}

	// Access the deserialized data
	fmt.Println("Name:", newPerson.Name)
	fmt.Println("ID:", newPerson.Id)
	fmt.Println("Email:", newPerson.Email)
}
