#!/bin/bash
cd ../docker

  #NUMWORKER
  echo "Numero di worker: "
  read -r NUMWORKER

  #NUMCLUSTER
  echo "Numero di cluster: "
  read -r NUMCLUSTER

  echo NUMWORKER="${NUMWORKER}"> ../.env
  echo NUMCLUSTER="${NUMCLUSTER}">> ../.env

echo NUMWORKER="${NUMWORKER}"

# Docker
docker compose --profile benchmark build
sleep 10
docker compose --profile app up benchmark_s --scale worker_s="${NUMWORKER}"
docker container start master
docker cp master:/doc ./