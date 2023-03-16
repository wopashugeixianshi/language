#!/bin/bash
for num in $(seq 1 500)
do
	docker rmi nginx:v$num
done
