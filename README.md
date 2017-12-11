Rosetta
=======

Rosetta is a simple, scriptable file encryption tool.

It works with symmetric keys, including a helpful password-derived key generator.

Files are saved in PEM encoding.  You know, that thing that has
"-----BEGIN FOOBAR-----" headers, and a bunch of Base64.
It looks like plaintext and is easy to copy-paste around.

The encryption is deterministic from the content.

Because of the deterministic encryption, base64 ascii armor, and scriptable CLI,
Rosetta is a great fit for keeping secrets in a git repo, encrypted.
It's easy to use both in local development and can be scripted for use in CI.



Development Status
------------------

Alpha.  As with anything computerful, use at your own risk.

Most of the CLI is still coming.  Hang on.

- Ciphers suites:
  - [nacl](https://godoc.org/golang.org/x/crypto/nacl/secretbox): ☑ supported.
  - anything else: PRs welcome.
- Password-derived keys: TODO
- CLI:
  - operations...
    - encrypt/decrypt: TODO
    - encryptFile/decryptFile: TODO
    - keyderive (standalone): TODO
    - bulk file mode: TODO
  - getting keys from...
    - password args: TODO
	- password env: TODO
	- key args: TODO
	- key env: TODO
	- key filename in args: TODO
	- key filename in env: TODO



CLI
---

The Rosetta CLI is straightforward:

  - `rosetta encrypt -k <keyfile> [-s <suite>]` -- streaming/pipe mode
  - `rosetta encrypt -p <password> [-s <suite>]` -- streaming/pipe mode
  - `rosetta encryptFile -k <keyfile> [-s <suite>] <cleartextPath> <ciphertextPath>`
  - `rosetta encryptFile -p <password> [-s <suite>] <cleartextPath> <ciphertextPath>`
  - `rosetta decrypt -k <keyfile> [-s <suite>]` -- streaming/pipe mode
  - `rosetta decrypt -p <password> [-s <suite>]` -- streaming/pipe mode
  - `rosetta decryptFile -k <keyfile> [-s <suite>] <ciphertextPath> <cleartextPath>`
  - `rosetta decryptFile -p <password> [-s <suite>] <ciphertextPath> <cleartextPath>`
  - `rosetta keyderive -p <password> <keyfile>` -- use to do the PBE just once;
    this is useful to save time for scripts that are about to do a ton of
    individual operations, because `-k` mode is generally ~100ms faster than `-p`.

The `encryptFile` commands are interchangable with
`cat cleartextPath | rosetta encrypt > ciphertextPath`, or even
`rosetta encrypt < cleartextPath > ciphertextPath` if you'd like to avoid `cat`,
and the same is correspondingly true for `decryptFile` and `decrypt`.

(TODO) a mode for operating on many files at once might also be useful:

  - `rosetta encryptFilesBulk -p <password> [-s <suite>]
    --clearBase <cleartextBasePath> --cipherBase <ciphertextBasePath>
    <files>...` -- you can clearly do this yourself in a script, but
    this will automate doing the keyfile derive once, prettyprint each
    operation, etc.
  - `rosetta encryptFilesBulk -k <keyfile> ...` -- as above, with keyfile.

Further TODO support:

  - any of the above `-p` modes should also support a `--envpw` flag,
    which switches it to reading an env var for the password.
  - we should do that for the key, too, so you don't *have* to
    flush that to disk in order to cache a key derivation.



Cipher choices
--------------

Encryption is using the well-standardized `nacl` system (specifically,
"secretbox" -- the symmetric key mode).  `nacl` is composed of the
XSalsa20 and Poly1305 primitives, ensuring both privacy and that the
ciphertext cannot be modified without holding the key.

If using passwords to derive keys, the well-standardized `scrypt` system
is used to generate a strong key.



Storage format
--------------

PEM -- RFC 1421 -- is Rosetta's storage format of choice.
PEM is probably already familiar to you as the format you see in encrypted
SSH keys, or TLS certificates, and other similar applications.



Caveats
-------

### file size

There is no strict limit to the size of file that can be encrypted with
Rosetta, but it's not designed with multi-gigabyte ISOs in mind either.
Rosetta may allocate `O(n)` memory proportional to the size of your file
when operating on it (both `nacl` and the PEM libraries we use require
the full content to be in memory at once), and since our storage format
involves Base64 encoding, the ciphertext file will be somewhat larger
than the original cleartext file.

But for any reasonably sized human-written document, yeah, chill out.
It's fine.

### security

You are responsible for the security of the device on which you run this program.

Fundamentally, if you decrypt a secret on an untrusted device, you lose.
This is true no matter what encryption standards or software you use.

There is no attempt to lock the passphrase into memory.  The passphrase may
be paged to disk or included in a core dump (should the program crash).

There is no attempt to keep you from saving keyfiles on disk.  If disk
archeology is part of your threat model, don't use those features.

There is no attempt to keep you from using "12345" as a password.

Rosetta trusts you to do the right thing.
