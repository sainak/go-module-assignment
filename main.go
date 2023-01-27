package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-faker/faker/v4"
)

type User struct {
	Name           string `faker:"name"`
	Age            uint   `faker:"boundary_start=25, boundary_end=45"`
	Email          string `faker:"email"`
	Phone          string `faker:"e_164_phone_number"`
	CreditCard     string `faker:"cc_number"`
	CreditCardType string `faker:"cc_type"`
}

// make user implement stringer
func (u User) String() string {
	b, err := json.MarshalIndent(u, "", "\t")
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return string(b)
}

func main() {
	var user User
	err := faker.FakeData(&user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user)
}
