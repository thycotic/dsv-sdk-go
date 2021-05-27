package vault

import (
	"fmt"
	"testing"
)

const secretName = "/test/secret"

// TestSecret tests Secret
func TestSecret(t *testing.T) {

	var err error
	secret := &Secret{
		secretResource: secretResource{
			Attributes: map[string]interface{}{"Test": "Attribute"},
			Data:       map[string]interface{}{"Secret": "Data"},
			Path:       secretName,
		},
		vault: *dsv,
	}

	t.Run("TestCreateSecret", func(t *testing.T) {
		err = dsv.NewSecret(secret)
		if err != nil {
			fmt.Println(err.Error())
			t.Error("creating secret:", err)
			return
		}
	})
	t.Run("TestUpdateSecret", func(t *testing.T) {
		secret.Attributes = map[string]interface{}{}
		secret.Data["Secret"] = "NewData"

		err = secret.Update(true)

		if err != nil {
			t.Error("updating secret:", err)
			return
		}
		if secret.Data == nil {
			t.Error("secret.Data is nil")
			return
		}
		if len(secret.Attributes) != 0 {
			t.Error("secret.Attributes not updated")
			return
		}
		if secret.Data["Secret"] != "NewData" {
			t.Error("secret.Data not updated")
		}
	})
	t.Run("TestGetSecret", func(t *testing.T) {
		secret, err = dsv.Secret(secretName)

		if err != nil {
			t.Error("calling secrets.Secret:", err)
			return
		}
		if secret.Data == nil {
			t.Error("secret.Data is nil")
			return
		}
		if len(secret.Attributes) != 0 {
			t.Error("secret.Attributes not updated")
			return
		}
		if secret.Data["Secret"] != "NewData" {
			t.Error("secret.Data not updated")
		}
	})
	t.Run("TestDeleteSecret", func(t *testing.T) {
		err = secret.Delete(true)
		if err != nil {
			t.Error("deleting secret:", err)
		}
	})

}
