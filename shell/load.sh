#!/bin/bash

mode=$(cat ./.mode)
sendfile=""

if test "$mode" = "go" ; then
  sendfile="main.go"
fi

if test "$mode" = "cpp"; then
  sendfile="main.cpp"
fi

dir=$(cat ./contest)
echo -e CONTEST: $dir
echo -n QUESTION?
read q
cp -i _result/_$dir/$q/$sendfile $sendfile
