package nacl

import (
	"fmt"

	"golang.org/x/crypto/nacl/secretbox"
	"golang.org/x/crypto/sha3"

	"github.com/polydawn/rosetta/cipher"
)

var _ cipher.EncryptBytesTool = EncryptBytes

func EncryptBytes(
	cleartext cipher.Cleartext, key cipher.Key,
) (
	ciphertext cipher.Ciphertext, nonce cipher.Nonce, err error,
) {
	// Format the key: nacl wants a 32 byte key, but our type is not so
	//  specific, so we must check and flip that.
	k2 := [32]byte{}
	if len(key) != 32 {
		return nil, nil, fmt.Errorf("invalid key: nacl keys must be 32 bytes (not %d)", len(key))
	}
	copy(k2[:], key)

	// Fabricate a nonce.
	nonce = deriveNonce(cleartext, key)
	n2 := [24]byte{}
	copy(n2[:], nonce)

	// Misc setup.
	prefix := []byte{} // nacl interface supports a prefix for some reason, but we don't use this.

	// Run cipher!
	ciphertext = secretbox.Seal(prefix, cleartext, &n2, &k2)
	return ciphertext, nonce, nil
}

// Derive a nonce deterministically from the cleartext and key, using an HMAC
// construction.  This is deterministic, but gives away no information about
// the cleartext unless you hold the key, aside from obviously revealing when
// the same cleartext is encrypted under the same key more than once (the
// same nonce, and thereafter the same ciphertext, will result).
func deriveNonce(cleartext cipher.Cleartext, key cipher.Key) cipher.Nonce {
	// This is implemented using the Keccak construction,
	// because Poly1305 documents itself as being unsafe to use twice with
	// the same key, which is a heck of a limitation.
	//
	// The requirement of nonces in Poly1305 is documented in the first
	// paragraph of the "Design decisions" section (page 6) of the paper:
	// https://cr.yp.to/mac/poly1305-20050329.pdf
	//
	// The reasoning there isn't exactly *unsound* per se, but it does seem
	// to make Poly1305 fairly impossible to use here.
	//
	// The use of Keccak here eschews HMAC wrapping because it can:
	// since Keccak is not subject to length-extension, this is a
	// well-documented and valid way to use it for a MAC-like role.
	// https://godoc.org/golang.org/x/crypto/sha3#example-package--Mac
	// We use the 128 variant rather than 256 since we'll be trucating
	// our output down to 24 bytes for the nacl hash anyway.
	nonce := make([]byte, 24)
	d := sha3.NewShake128()
	d.Write(key)
	d.Write(cleartext)
	d.Read(nonce)
	return nonce
}

var _ cipher.DecryptBytesTool = DecryptBytes

func DecryptBytes(
	ciphertext cipher.Ciphertext, key cipher.Key, nonce cipher.Nonce,
) (
	cleartext cipher.Cleartext, err error,
) {
	// Format the key: nacl wants a 32 byte key, but our type is not so
	//  specific, so we must check and flip that.
	k2 := [32]byte{}
	if len(key) != 32 {
		return nil, fmt.Errorf("invalid key: nacl keys must be 32 bytes (not %d)", len(key))
	}
	copy(k2[:], key)

	// Format the nonce: nacl wants a 24 byte nonce, but our type is not so
	//  specific, so we must check and flip that.
	n2 := [24]byte{}
	if len(nonce) != 24 {
		return nil, fmt.Errorf("invalid nonce: nacl nonce must be 24 bytes (not %d)", len(key))
	}
	copy(n2[:], nonce)

	// Misc setup.
	reuseMemory := []byte{} // nacl interface supports reusing existing byte slices, but we don't use this.

	// Run cipher!
	cleartext, success := secretbox.Open(reuseMemory, ciphertext, &n2, &k2)
	if !success {
		return nil, fmt.Errorf("decryption failed") // no reason!  :/
	}
	return cleartext, nil
}
