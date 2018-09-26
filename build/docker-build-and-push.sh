#!/bin/bash

echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin

array=( frontend jobs-api worker )
for i in "${array[@]}"
do
  echo "Building docker image for $i"
	docker build -t $i ./$i
  docker tag $i alexandrevilain/image-resizer-$i
  docker push alexandrevilain/image-resizer-$i
done