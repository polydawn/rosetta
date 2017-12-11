package rosetta

import (
	"strings"
	"testing"

	"github.com/pmezard/go-difflib/difflib"
)

func TestEnvelope(t *testing.T) {
	fixtureStruct := struct {
		body    []byte
		headers map[string]string
	}{
		[]byte("blah blah blah"),
		map[string]string{"k": "v", "a": "b"},
	}
	fixtureEnvelope := strings.Trim(strings.Replace(`
		-----BEGIN ROSETTA CIPHERTEXT-----
		a: b
		k: v
		
		YmxhaCBibGFoIGJsYWg=
		-----END ROSETTA CIPHERTEXT-----
	`, "\t", "", -1), "\n") + "\n"

	t.Run("envelope", func(t *testing.T) {
		raw, err := EnvelopeBytes(fixtureStruct.body, fixtureStruct.headers)
		if err != nil {
			t.Fatalf("%s", err)
		}
		wantStringEqual(t, fixtureEnvelope, string(raw))
	})
	t.Run("unenvelope", func(t *testing.T) {
		body, headers, err := UnenvelopeBytes([]byte(fixtureEnvelope))
		if err != nil {
			t.Fatalf("%s", err)
		}
		for headerKey, headerValue := range fixtureStruct.headers {
			if headers[headerKey] != headerValue {
				t.Fatalf("lost header %q: expect %q, got %q", headerKey, headerValue, headers["a"])
			}
		}
		wantStringEqual(t, string(fixtureStruct.body), string(body))
	})
}

func wantStringEqual(t *testing.T, a, b string) {
	t.Helper()
	result, err := difflib.GetUnifiedDiffString(difflib.UnifiedDiff{
		A:       difflib.SplitLines(a),
		B:       difflib.SplitLines(b),
		Context: 3,
	})
	if err != nil {
		t.Fatalf("diffing failed: %s", err)
	}
	if result != "" {
		t.Errorf("Match failed: diff:\n%s", result)
	}
}
