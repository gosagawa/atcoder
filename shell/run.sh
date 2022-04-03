#!/bin/bash

mode=$(cat ./.mode)

if test "$mode" = "go" ; then
  go run main.go i
fi

if test "$mode" = "cpp"; then
  g++ -D=__LOCAL -o main main.cpp
  ./main
fi
