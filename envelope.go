package rosetta

import (
	"github.com/polydawn/rosetta/cipher"
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
	return nil, nil
}

func UnenvelopeAndDecryptBytes(raw []byte, key cipher.Key) (cipher.Cleartext, error) {
	return nil, nil
}

func EnvelopeBytes(ciphertext cipher.Ciphertext, headers map[string]string) ([]byte, error) {
	return nil, nil
}

func UnenvelopeBytes(raw []byte) (body cipher.Ciphertext, headers map[string]string, err error) {
	return nil, nil, nil
}

func EncryptBytes(
	cleartext cipher.Cleartext, key cipher.Key,
) (
	ciphertext cipher.Ciphertext, nonce cipher.Nonce, err error,
) {
	return nil, nil, nil
}

func DecryptBytes(
	ciphertext cipher.Ciphertext, key cipher.Key, nonce cipher.Nonce,
) (
	cleartext cipher.Cleartext, err error,
) {
	return nil, nil
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
