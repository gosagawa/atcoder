#!/bin/bash

mode=$(cat ./.mode)
echo -n URL?
read q

if test "$mode" = "go" ; then
  oj s $q main.go 
fi

if test "$mode" = "cpp"; then
  oj s $q main.cpp 
fi
