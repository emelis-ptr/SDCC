#!/bin/bash
cd ../docker

echo "Benchmark after execution main [y] - Benchmark with no execution main[n]"
read benchmark

case $benchmark in
y) echo ""
  #NUMWORKER
    echo "Inserisci numero worker: "
    read -r NUMWORKER
  ;;
n) echo ""
  #NUMWORKER
  echo "Numero di worker: "
  read -r NUMWORKER

  #NUMCLUSTER
  echo "Numero di cluster: "
  read -r NUMCLUSTER

  echo NUMWORKER="${NUMWORKER}"> ../.env
  echo NUMCLUSTER="${NUMCLUSTER}">> ../.env
  ;;
?)
  ;;
esac

echo NUMWORKER="${NUMWORKER}"

# Docker
docker compose --profile benchmark build
sleep 10
docker compose --profile app up benchmark_s --scale worker_s="${NUMWORKER}"