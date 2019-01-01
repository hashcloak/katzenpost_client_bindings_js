
.. image:: https://godoc.org/github.com/katzenpost/client_bindings?status.svg
  :target: https://godoc.org/github.com/katzenpost/client_bindings

C language bindings for Client
==============================

Here we provide C language bindings for the mixnet client library:

* https://github.com/katzenpost/client


Building
--------

Firstly, acquire a recent version of dep: https://github.com/golang/dep

And then fetch the pinned dependencies:
::
   dep ensure

We generate our bindings.h and bindings.so files like so:
::
   go build -o client_bindings.so -buildmode=c-shared bindings.go


The provided example can then be built with:
::
   gcc ./examples/ping.c ./client_bindings.so -I . -o cping


license
=======

AGPL: see LICENSE file for details.
