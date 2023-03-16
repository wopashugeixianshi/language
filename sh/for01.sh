#!/bin/bash

for num in $(seq 400 500)
do
#echo $num
#docker tag nginx:latest nginx:v$num
docker rmi nginx:v$num
done
