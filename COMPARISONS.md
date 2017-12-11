Things that aren't Rosetta
==========================

A variety of other programs already use the `nacl` golang implementation, and
since searching for users of those packages is fairly easy, we did so before
starting to write Rosetta.  At the time of writing, several things Looks
*close* to the feature set Rosetta aims for, but none were quite there:

- https://github.com/alistanis/goenc
  - has no CLI
- https://github.com/danderson/gobox
  - ALMOST, but supports only interactive passwords, and it's with that hackery that ssh uses to stop you from piping it in either... and I actually need/want to use this headless, in CI, so, nope.
- https://github.com/andmarios/golang-nacl-secretbox
  - ... there are a *lot* of half-baked things floating around on the internet.
- https://github.com/avahowell/masterkey
  - cool, but built to be a password manager.  I want something simpler that's meant to be scripted around.
- https://github.com/meltwater/secretary
  - cool, but again, built to interact with way to much -- I don't use AWS KMS nor Mesos; I want something simpler that's meant to be scripted around.
- https://github.com/nutmegdevelopment/nutcracker
  - cool, but again, I don't want a REST API, I just want to script around something and pipe some dang streams.
- https://github.com/scode/saltybox
  - ALMOST, but like danderson/gobox, supports only interactive passwords, and I actually need/want to use this headless, in CI, so, nope.
- https://github.com/tsileo/blobs-cli
  - ... doesn't even have a readme?  Looks almost relevant, but no docs, and also seems to be content aware, so nope and nope.
- https://github.com/voutasaurus/box
  - ALMOST ALMOST!!  I could see myself using this.  The main improvements Rosetta offers over this is: more key gen options (particularly, password-derived keys), and the PEM format for ciphertext.
  - Oh, also it doesn't do deterministic encryption, because it picks a fresh salt every time.  Also also it *does that wrong*, picking from `math/rand` rather than secure random.  Not that that's cataclysmic for a salt, but, plz.
- https://github.com/xrstf/boxer
  - has no CLI.  the nacl API is not complex enough that another wrapper library is necessary.
- https://github.com/justwatchcom/gopass
  - This is an extremely mature project, unlike most of the other listed, and probably great... if want you want is a password manager.
  - As previously commented: we want deterministic encryption, and we want simple scripting, and non-interactive use.  Gopass scores partial points on only some of these requirements.
  - `gopass` exec's out to the system `gpg` command, which is almost as irritating of an API to version against or ship itself as `openssl`.  Rosetta does not and will not depend on such things.  Rosetta is pure go and easy to cross compile and ship anywhere.
