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

You are responsible for the security of the device on which you run this program.

Fundamentally, if you decrypt a secret on an untrusted device, you lose.
This is true no matter what encryption standards or software you use.

There is no attempt to lock the passphrase into memory.  The passphrase may
be paged to disk or included in a core dump (should the program crash).

There is no attempt to keep you from saving keyfiles on disk.  If disk
archeology is part of your threat model, don't use those features.

There is no attempt to keep you from using "12345" as a password.

Rosetta trusts you to do the right thing.
