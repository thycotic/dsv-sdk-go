package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/thycotic/dsv-sdk-go/pkg/vault"
)

func main() {
	secretPath := "/test/secret"

	flag.StringVar(&secretPath, "secret", secretPath, "the path of the secret")
	flag.Parse()

	var dsv *vault.Vault

	if config, err := vault.GetConfigFromEnv(); err != nil {
		log.Fatalf("[ERROR] DSV_SDK_CONFIG or DSV_SDK_CONFIG_File must be set")
	} else if config != nil {
		if dsv, err = vault.New(*config); err != nil {
			log.Fatalf("configuration error: %s", err)
		}
	}

	if secret, err := dsv.Secret(secretPath); err != nil {
		log.Fatalf("[ERROR] getting secret: %s", err)
	} else {
		fmt.Println(secret.Data)
	}
}
