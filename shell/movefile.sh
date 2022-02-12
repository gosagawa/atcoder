#!/bin/bash
dir=$(cat ./contest)
echo -e CONTEST: $dir
echo -n QUESTION?
read q
mkdir -p _result/_$dir/$q
mv -i $1 _result/_$dir/$q
git add _result/_$dir/$q/$1
git commit -m "$dir $q"
cp -i _template/$1 ./$1

