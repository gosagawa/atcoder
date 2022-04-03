#!/bin/bash

mode=$(cat ./.mode)
sendfile=""

if test "$mode" = "go" ; then
  oj t -c "go run main.go"
fi

if test "$mode" = "cpp"; then
  g++ main.cpp; oj t 
fi
