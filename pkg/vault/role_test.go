package vault

import (
	"testing"
)

const (
	notFound = "404 Not Found: {\"code\":404,\"message\":\"unable to find item with specified identifier\"}"
	roleName = "test-role"
)

// TestRole tests Role
func TestRole(t *testing.T) {
	config, _ := GetConfigFromEnv()
	vault, _ := New(*config)

	t.Run("Get", func(t *testing.T) {
		if role, err := vault.Role(roleName); err != nil {
			t.Errorf("calling roles.Role(\"%s\"): %s", roleName, err)
		} else if role.Name != roleName {
			t.Errorf("expecting %s but roles.Role was %s", roleName, role.Name)
		}
	})
	t.Run("GetNonexistent", func(t *testing.T) {
		roleName := "nonexistent-role"
		if role, err := vault.Role(roleName); err == nil {
			t.Errorf("role '%s' exists but but it should not", roleName)
		} else if err != nil && err.Error() != notFound {
			t.Errorf("unexpected error calling roles.Role(\"%s\"): %s", roleName, err)
		} else if role != nil {
			t.Errorf("roles.Role returned a role and no error for '%s'", roleName)
		}
	})
}
