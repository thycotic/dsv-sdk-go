package vault

import "testing"

const secretName = "/test/secret"

// TestSecret tests Secret
func TestSecret(t *testing.T) {
	config, _ := GetConfigFromEnv()
	vault, _ := New(*config)

	if secret, err := vault.Secret(secretName); err != nil {
		t.Error("calling secrets.Secret:", err)
	} else if secret.Data == nil {
		t.Error("secret.Data is nil")
	}
}
