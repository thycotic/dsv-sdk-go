package main

import (
	"fmt"
	"log"
	"os"

	"github.com/thycotic/dsv-sdk-go/auth"
	"github.com/thycotic/dsv-sdk-go/vault"
)

func main() {
	dsv, err := vault.New(vault.Configuration{
		Tenant:   os.Getenv("DSV_TENANT"),
		Provider: auth.AWS,
	})

	if err != nil {
		log.Fatalf("failed to configure vault: %v", err)
	}

	secret, err := dsv.Secret("resources:us-east-5:server1")

	if err != nil {
		log.Fatalf("failed to fetch secret: %v", err)
	}

	fmt.Printf("secret data: %v", secret.Data)
	fmt.Printf("secret attributes: %v", secret.Attributes)
}
