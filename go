#!/bin/bash
set -eux

[ -e configure ] || ./autogen.sh
[ -e Makefile ] || ./configure
make -j$(nproc)

sudo rm -vfr \
	/usr/local/include/pixman* \
	/usr/local/lib/libpixman* \
	/usr/local/lib/pkgconfig/pixman*

sudo make install

sudo ldconfig /usr/local/lib
