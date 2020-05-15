package vault

import (
	"encoding/json"
	"log"
	"net/url"
	"strconv"
)

// secretsResource is the HTTP URL path component for the secrets resource
const secretsResource = "secrets"

// secretUpdateBody is the Secret composed with update specific fields
type secretUpdateBody struct {
	Secret
	Overwrite bool
}

// secretResource is composed with resourceMetadata to for SecretContents
type secretResource struct {
	Attributes map[string]interface{}
	Data       map[string]interface{}
	Path       string
}

// Secret holds the contents of a secret from DSV
type Secret struct {
	resourceMetadata
	secretResource
	vault Vault
}

// Secret gets the secret at path from the DSV of the given tenant
func (v Vault) Secret(path string) (*Secret, error) {
	secret := &Secret{vault: v}
	data, err := v.accessResource("GET", secretsResource, path, nil)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, secret); err != nil {
		log.Printf("[DEBUG] error parsing response from /%s/%s: %q", secretsResource, path, data)
		return nil, err
	}
	return secret, nil
}

// Delete deletes the secret from the DSV. Flag can be provided to force deletion
func (s Secret) Delete(force bool) error {
	params := url.URL{
		RawQuery: url.Values{
			"force": []string{strconv.FormatBool(force)},
		}.Encode(),
	}
	if _, err := s.vault.accessResource("DELETE", secretsResource, s.Path+params.String(), nil); err != nil {
		return err
	}

	return nil
}

// Update updates the secret in DSV
func (s Secret) Update(overwrite bool) error {
	// Force API to update empty Attributes
	if s.Attributes != nil && len(s.Attributes) == 0 {
		s.Attributes = map[string]interface{}{"": ""}
		defer func() { s.Attributes = map[string]interface{}{} }()
	}

	input := secretUpdateBody{
		Secret:    s,
		Overwrite: overwrite,
	}

	data, err := s.vault.accessResource("PUT", secretsResource, s.Path, input)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &s); err != nil {
		log.Printf("[DEBUG] error parsing response from /%s: %q", secretsResource, data)
		return err
	}

	return nil
}

// NewSecret creates a new Secret with given data
func (v Vault) NewSecret(secret *Secret) error {
	data, err := v.accessResource("POST", secretsResource, secret.Path, secret)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, secret); err != nil {
		log.Printf("[DEBUG] error parsing response from /%s: %q", secretsResource, data)
		return err
	}

	return nil
}
