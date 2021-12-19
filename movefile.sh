#!/bin/bash
dir=$(cat ./contest)
echo -e DIR: $dir
echo -n QUESTION?
read q
mkdir -p _result/_$dir/$q
mv -i main.go _result/_$dir/$q
git add _result/_$dir/$q/main.go
git commit -m "$dir $q"
cp -i _template/main.go ./main.go

