package vault

import "testing"

func TestClient(t *testing.T) {
	config, _ := GetConfigFromEnv()
	vault, _ := New(*config)

	var id string // set by TestNewClient used by Get and removed by Delete

	t.Run("New", func(t *testing.T) {
		client := &Client{clientResource: clientResource{RoleName: roleName}}

		if err := vault.New(client); err != nil {
			t.Errorf("calling clients.New(\"%s\"): %s", roleName, err)
		} else if client.ClientID == "" {
			t.Error("contents.ClientID was empty")
		} else if client.ClientSecret == "" {
			t.Error("contents.ClientSecret was empty")
		} else {
			id = client.ClientID
		}
		return
	})
	t.Run("Get", func(t *testing.T) {
		if client, err := vault.Client(id); err != nil {
			t.Errorf("calling clients.Client(\"%s\"): %s", id, err)
		} else if client.ClientID != id {
			t.Errorf("expecting %s but clients.Client was %s", id, config.Credentials.ClientID)
		}
	})
	t.Run("Delete", func(t *testing.T) {
		if client, err := vault.Client(id); err != nil {
			t.Errorf("calling clients.Client(\"%s\"): %s", id, err)
		} else if err := client.Delete(); err != nil {
			t.Errorf("calling client.Delete on Client %s: %s", id, err)
		}
	})
}
