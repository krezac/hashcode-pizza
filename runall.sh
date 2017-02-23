#!/bin/sh

rm -f hascode-pizza
rm -f data/*.out
go build
./hashcode-pizza data/example.in
./hashcode-pizza data/small.in
./hashcode-pizza data/medium.in
./hashcode-pizza data/big.in
