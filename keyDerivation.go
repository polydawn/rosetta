package rosetta

import (
	"golang.org/x/crypto/scrypt"
	"golang.org/x/crypto/sha3"

	"github.com/polydawn/rosetta/cipher"
)

func DeriveKey(password []byte, keySize int) (cipher.Key, error) {
	salt := deriveSalt(password)
	// The constants for size and difficulty here are the defaults
	// recommending in the docs: "The recommended parameters for interactive
	// logins as of 2017 are N=32768, r=8 and p=1".
	// See https://godoc.org/golang.org/x/crypto/scrypt#Key
	key, err := scrypt.Key(password, salt, 1<<15, 8, 1, keySize)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// As is typical in Rosetta, we'll be making the salt, like our nonces,
// deterministic on the user input.  This fulfills all of the hardening
// purposes of a salt, without requiring a salt state to be kept around.
func deriveSalt(input []byte) cipher.Nonce {
	// Keccak construction is used to do the heavy lifting again.
	// This is not a keyed MAC, per se, though it looks similar.
	// The purpose of the prefix bytes is not any critical security purpose;
	// just a "why not" sort of shift of the space.
	nonce := make([]byte, 24)
	d := sha3.NewShake128()
	d.Write([]byte{'r', 'o', 's', 'e', 't', 't', 'a'})
	d.Write(input)
	d.Read(nonce)
	return nonce
}
