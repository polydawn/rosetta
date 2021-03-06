package rosetta

import (
	"bytes"
	"encoding/base64"
	"encoding/pem"
	"fmt"

	"github.com/polydawn/rosetta/cipher"
	"github.com/polydawn/rosetta/cipher/impl/nacl"
)

func EncryptFile(
	cleartextPath string, // from
	ciphertextPath string, // to
	key cipher.Key,
) error {
	return nil
}

func DecryptFile(
	ciphertextPath string, // from
	cleartextPath string, // to
	key cipher.Key,
) error {
	return nil
}

func EncryptAndEnvelopeBytes(cleartext cipher.Cleartext, key cipher.Key) ([]byte, error) {
	ciphertext, nonce, err := EncryptBytes(cleartext, key)
	if err != nil {
		return nil, err
	}
	headers := map[string]string{
		"cipher": "nacl",
		"nonce":  base64.StdEncoding.EncodeToString(nonce),
	}
	return EnvelopeBytes(ciphertext, headers)
}

func UnenvelopeAndDecryptBytes(raw []byte, key cipher.Key) (cipher.Cleartext, error) {
	ciphertext, headers, err := UnenvelopeBytes(raw)
	if err != nil {
		return nil, err
	}
	if cipherHeader, ok := headers["cipher"]; ok {
		switch cipherHeader {
		case "nacl":
			// good.  also, only one we have, so no logic here.
		default:
			return nil, fmt.Errorf("invalid envelope: unsupported cipher %q", cipherHeader)
		}
	} else {
		return nil, fmt.Errorf("invalid envelope: missing required header %q", "cipher")
	}
	var nonce cipher.Nonce
	if nonceHeader, ok := headers["nonce"]; ok {
		nonceStr, err := base64.StdEncoding.DecodeString(nonceHeader)
		if err != nil {
			return nil, fmt.Errorf("error opening envelope: could not decode nonce")
		}
		nonce = []byte(nonceStr)
	}
	return DecryptBytes(ciphertext, key, nonce)
}

func EnvelopeBytes(ciphertext cipher.Ciphertext, headers map[string]string) ([]byte, error) {
	block := pem.Block{
		Type:    "ROSETTA CIPHERTEXT",
		Headers: headers,
		Bytes:   ciphertext,
	}
	var buf bytes.Buffer
	if err := pem.Encode(&buf, &block); err != nil {
		panic(err) // we're writing to an in-memory buffer... not much can go wrong
	}
	return buf.Bytes(), nil
}

func UnenvelopeBytes(raw []byte) (body cipher.Ciphertext, headers map[string]string, err error) {
	block, rest := pem.Decode(raw)
	if block == nil {
		return nil, nil, fmt.Errorf("invalid envelope: this doesn't look like a ciphertext at all! no envelope header found")
	}
	_ = rest // TODO check for non-whitespace leftovers... those actually indicate error
	if block.Type == "" {
		return nil, nil, fmt.Errorf("invalid envelope: this doesn't look like a ciphertext at all")
	}
	if block.Type != "ROSETTA CIPHERTEXT" {
		return nil, nil, fmt.Errorf("invalid envelope: this doesn't look like a rosetta ciphertext")
	}
	return block.Bytes, block.Headers, nil
}

func EncryptBytes(
	cleartext cipher.Cleartext, key cipher.Key,
) (
	ciphertext cipher.Ciphertext, nonce cipher.Nonce, err error,
) {
	return nacl.EncryptBytes(cleartext, key)
}

func DecryptBytes(
	ciphertext cipher.Ciphertext, key cipher.Key, nonce cipher.Nonce,
) (
	cleartext cipher.Cleartext, err error,
) {
	return nacl.DecryptBytes(ciphertext, key, nonce)
}

/*
	Rosetta stores the encrypted seralized data in a file of the familar
	ascii-armored format, with PEM-style "-----BEGIN WHATSIT-----" and
	"-----END WHOSIT-----" prefixes and suffixes.

	This format helps make Rosetta encrypted files declare who they are
	and how to work with them (or at least how to google how to work with
	them!), and the ascii-armor makes them less of pain to render and diff
	(even if the diff is useless) in lots of places where you might want to
	store and transport them -- for example in git.
*/
type Envelope struct {
}
