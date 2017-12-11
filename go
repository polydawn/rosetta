#!/bin/bash
PKG="github.com/polydawn/rosetta"

# Always ensure cwd is the script dir / repo root, for simplicity.
cd "$( dirname "${BASH_SOURCE[0]}" )"

# Export local gopath, mkdir it, make self-symlink for globally consistent imports, and go there.
export GOBIN="$PWD/bin"
export GOPATH="$PWD/.gopath/"
mkdir -p "$(dirname "$GOPATH/src/$PKG")"
ln -snf "$(echo "${PKG//[^\/]}/" | sed s#/#../#g)"../ "$GOPATH/src/$PKG"
cd "$GOPATH/src/$PKG"
# Use ALL GO COMMANDS as NORMAL.  You now have a project-local scoped gopath!  NBD.
go "$@"
