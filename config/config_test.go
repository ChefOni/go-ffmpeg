package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	ValidKey   = "CLOUDFLARE_API_KEY"
)

func TestLoadStorageBucketDetails(t *testing.T) {
	t.Run("Valid key", func(t *testing.T) {
		t.Setenv(ValidKey, "fjfjfpncri3n03fuo3f93v39u")
		value, err := LoadStorageBucketDetails(ValidKey)
		assert.NoError(t, err)
		assert.Equal(t, "fjfjfpncri3n03fuo3f93v39u", value)
	})

	t.Run("Key with no value", func(t *testing.T) {
		t.Setenv("EMPTY_KEY", "")
		value, err := LoadStorageBucketDetails("EMPTY_KEY")
		assert.Error(t, err)
		assert.Empty(t, value)
		assert.Contains(t, err.Error(), "EMPTY_KEY")
	})
}
