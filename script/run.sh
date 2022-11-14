#!/bin/bash
cd ..

echo "Numero di worker: "
read -r NUMWORKER

echo "Numero di punti: "
read -r NUMPOINT

echo "Numero di cluster: "
read -r NUMCLUSTER

echo "Numero di mapper: "
read -r NUMMAPPER

echo "Numero di reducer: "
read -r NUMREDUCER

if [ "${NUMCLUSTER}" -ge "${NUMPOINT}" ]; then
  echo "Numero cluster maggiore dell'insieme dei punti. Riprova"
  read NUMCLUSTER
fi

ALGO=""

echo "Select [1]: LLyod, [2]: standard kmeans, [3]: keans plus plus"
read algos
case $algos in
1) echo "llyod"
  ALGO=llyod
  ;;
2) echo "standardKMeans"
  ALGO=standardKMeans
  ;;
3) echo "kmeans++"
  ALGO=kmeansAlgo
  ;;
esac

echo NUMWORKER="${NUMWORKER}" > .env
# shellcheck disable=SC2129
echo NUMPOINT="${NUMPOINT}">> .env
echo NUMCLUSTER="${NUMCLUSTER}">> .env
echo ALGO=${ALGO}>> .env
echo NUMMAPPER="${NUMMAPPER}">> .env
echo NUMREDUCER="${NUMREDUCER}">> .env

docker compose build

sleep 10

docker compose up --scale worker_s="${NUMWORKER}"