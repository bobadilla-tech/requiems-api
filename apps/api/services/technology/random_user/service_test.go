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
