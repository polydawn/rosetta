package rosetta

import (
	"encoding/base64"
	"testing"
)

// Note well: this test should take *around* a tenth of a second.
// It's *supposed* to be doing hard work.
func TestDeriveKey(t *testing.T) {
	keyFixture := "Imk/4h9LMsBDGfeZaGyIFBfUKQ+/KODS+j0RCH0Yp7A="
	key, err := DeriveKey([]byte("asdf"), 32)
	if err != nil {
		t.Errorf("%s", err)
	}
	if len(key) != 32 {
		t.Errorf("got key of length %d", len(key))
	}
	b64key := base64.StdEncoding.EncodeToString(key)
	if keyFixture != b64key {
		t.Errorf("key should be %q, got %q", keyFixture, b64key)
	}
}
