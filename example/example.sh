#!/bin/bash

if [ $# -ne 1 ]; then
	echo "please provide how many qr codes do you want to generate"
	exit 1
fi

SHA1=''

function randomSHA1() {
	SHA1=`echo "$RANDOM-with-salt" | sha1sum | awk '{print $1}'`
}


for i in $(seq 1 $1)
do
randomSHA1
echo $SHA1
done
