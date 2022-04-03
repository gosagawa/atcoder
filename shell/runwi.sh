#!/bin/bash

mode=$(cat ./.mode)

if test "$mode" = "go" ; then
  go run main.go
fi

if test "$mode" = "cpp"; then
  g++ -o main main.cpp
  ./main
fi
