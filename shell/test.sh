#!/bin/bash

mode=$(cat ./.mode)
sendfile=""

if test "$mode" = "go" ; then
  oj t -c "./main"
fi

if test "$mode" = "cpp"; then
  g++ main.cpp; oj t 
fi
