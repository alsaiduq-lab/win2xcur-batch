#!/usr/bin/env nix-shell
#!nix-shell -i bash -p go python3 uv imagemagick
set -euo pipefail

uv venv .venv
source .venv/bin/activate
uv pip install -qU win2xcur

binpath="$(command -v convert || command -v magick)"
store_root="$(readlink -f "$binpath")"
store_root="${store_root%/bin/*}"
export LD_LIBRARY_PATH="${store_root}/lib:${LD_LIBRARY_PATH-}"
[ -d "${store_root}/etc/ImageMagick-7" ] && export MAGICK_CONFIGURE_PATH="${store_root}/etc/ImageMagick-7"

go build
$PWD/win2xcur-batch
