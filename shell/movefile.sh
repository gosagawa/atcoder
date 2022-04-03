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
mkdir -p _result/_$dir/$q
mv -i $sendfile _result/_$dir/$q
git add _result/_$dir/$q/$sendfile
git commit -m "$dir $q"
cp -i _template/$sendfile ./$sendfile

