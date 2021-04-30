#!/bin/bash

echo -n DIR?
read dir
echo -n QUESTION?
read q
mkdir -p _result/_$dir/$q
mv -i main.go _result/_$dir/$q
git add _result/_$dir/$q/main.go
git commit -m "$dir $q"

