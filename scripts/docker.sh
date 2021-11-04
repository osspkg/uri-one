#!/bin/bash

cd $PWD

docker_up() {
  docker-compose -f deployments/docker-compose.yaml -p dev_uri-one up
}

docker_down() {
  docker-compose -f deployments/docker-compose.yaml -p dev_uri-one down
}

case $1 in
docker_up)
  docker_down
  docker_up
  ;;
docker_down)
  docker_down
  ;;
*)
  echo "docker_up or docker_down"
  ;;
esac
