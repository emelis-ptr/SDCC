#!/bin/bash
cd ..

echo "Numero di worker: "
read NUMWORKER

echo "Numero di punti: "
read NUMPOINT

echo "Numero di cluster: "
read NUMCLUSTER

if [ "${NUMCLUSTER}" -ge "${NUMPOINT}" ]; then
  echo "Numero cluster maggiore dell'insieme dei punti. Riprova"
  read NUMCLUSTER
fi

ALGO=""

echo "${NUMWORKER}"
echo "${NUMPOINT}"
echo "${NUMCLUSTER}"

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
# shellcheck disable=SC2086
echo NUMCLUSTER=${NUMCLUSTER}>> .env
echo ALGO=${ALGO}>> .env

docker compose build

sleep 10

docker compose up --scale worker_s=${NUMWORKER}