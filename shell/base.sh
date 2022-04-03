#!/bin/bash

mode=$(cat ./.mode)
sendfile=""

if test "$mode" = "go" ; then
  cp -i _template/main.go ./main.go
fi

if test "$mode" = "cpp"; then
  cp -i _template/main.cpp ./main.cpp
fi
