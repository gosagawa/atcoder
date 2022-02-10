#!/bin/bash
echo -n URL?
read q
rm -rf test
echo $q
oj download $q
cp test/sample-1.in input
