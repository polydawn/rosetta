Rosetta
=======

Rosetta is a simple, scriptable file encryption tool.

It works with symmetric keys, including a helpful password-derived key generator.

Files are saved in PEM encoding.  You know, that thing that has "-----BEGIN FOOBAR-----"
headers, and a bunch of Base64.

The encryption is deterministic from the content.

Because of the deterministic encryption, base64 ascii armor, and scriptable CLI,
Rosetta is a great fit for keeping secrets in a git repo, encrypted.
It's easy to use both in local development and can be scripted for use in CI.
