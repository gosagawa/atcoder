#!/bin/bash
mode=$(cat ./.mode)

if test "$mode" = "go" ; then
  find ./ -maxdepth 1  -name main.go -or -name input | entr -c ./shell/run.sh
fi

if test "$mode" = "cpp"; then
  find ./ -maxdepth 1  -name main.cpp -or -name input | entr -c ./shell/run.sh
fi


