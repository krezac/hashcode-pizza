#!/bin/sh

rm -f hascode-pizza
rm -f data/*.out
go build
#./hashcode-pizza data/kittens.in
./hashcode-pizza data/me_at_the_zoo.in
#./hashcode-pizza data/trending_today.in
#./hashcode-pizza data/videos_worth_spreading.in
