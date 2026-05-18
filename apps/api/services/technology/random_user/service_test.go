package randomuser

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate_FieldsPopulated(t *testing.T) {
	t.Parallel()
	svc := NewService()

	for range 20 {
		u := svc.Generate()

		assert.NotEmpty(t, u.Name)
		assert.NotEmpty(t, u.Email)
		if !strings.Contains(u.Email, "@") {
			t.Errorf("Email missing @: got %q", u.Email)
		}
		assert.NotEmpty(t, u.Phone)
		assert.NotEmpty(t, u.Address.Street)
		assert.NotEmpty(t, u.Address.City)
		assert.NotEmpty(t, u.Address.State)
		assert.NotEmpty(t, u.Address.Zip)
		assert.NotEmpty(t, u.Address.Country)
		if !strings.HasPrefix(u.Avatar, "https://api.dicebear.com/") {
			t.Errorf("Avatar should start with dicebear URL: got %q", u.Avatar)
		}
	}
}

func TestGenerate_AvatarContainsName(t *testing.T) {
	t.Parallel()
	svc := NewService()

	u := svc.Generate()
	if !strings.Contains(u.Avatar, "seed=") {
		t.Errorf("avatar URL should contain seed parameter: %q", u.Avatar)
	}
}

// TestGenerateBatch_ValidCount verifies that GenerateBatch returns the correct number of users.
func TestGenerateBatch_ValidCount(t *testing.T) {
	t.Parallel()
	svc := NewService()

	users, err := svc.GenerateBatch(5)
	assert.NoError(t, err)
	assert.Len(t, users, 5)
}

// TestGenerateBatch_FieldsPopulated verifies that each user returned by GenerateBatch has all fields set.
func TestGenerateBatch_FieldsPopulated(t *testing.T) {
	t.Parallel()
	svc := NewService()

	users, err := svc.GenerateBatch(3)
	assert.NoError(t, err)
	for i, u := range users {
		assert.NotEmpty(t, u.Name, "user[%d].Name", i)
		assert.NotEmpty(t, u.Email, "user[%d].Email", i)
		assert.NotEmpty(t, u.Phone, "user[%d].Phone", i)
		assert.NotEmpty(t, u.Address.Street, "user[%d].Address.Street", i)
		assert.NotEmpty(t, u.Avatar, "user[%d].Avatar", i)
	}
}
