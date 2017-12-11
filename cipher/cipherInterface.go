package cipher

type (
	Cleartext  []byte
	Ciphertext []byte
	Key        []byte
	Nonce      []byte
)

type (
	EncryptBytesTool func(Cleartext, Key) (Ciphertext, Nonce, error)
	DecryptBytesTool func(Ciphertext, Key, Nonce) (Cleartext, error)
)
