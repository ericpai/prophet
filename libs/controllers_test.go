package libs

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestCheckKeysExist(t *testing.T) {
	v := url.Values{
		"key_a": []string{"a"},
		"key_b": []string{"b"},
	}
	assert.True(t, CheckKeysExist(v, "key_a", "key_b"))
	assert.True(t, CheckKeysExist(v, "key_a"))
	assert.False(t, CheckKeysExist(v, "key_a", "key_b", "key_c"))
}
