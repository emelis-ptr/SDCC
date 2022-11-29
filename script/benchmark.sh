#!/bin/bash
cd ../docker

  #NUMWORKER
#NUMWORKER
echo "Numero di worker: "
read -r NUMWORKER

#NUMCLUSTER
echo "Numero di cluster: "
read -r NUMCLUSTER

if [ "${NUMCLUSTER}" -ge "${NUMPOINT}" ]; then
  echo "Numero cluster maggiore dell'insieme dei punti. Riprova"
  read NUMCLUSTER
fi

#ALGO
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
  ALGO=kmeansPlusPlus
  ;;
esac

# Write file
echo NUMWORKER="${NUMWORKER}" > ../.env
echo NUMCLUSTER="${NUMCLUSTER}">> ../.env
echo ALGO=${ALGO}>> ../.env

# Docker
docker compose --profile benchmark build
sleep 10
docker compose --profile app up benchmark_s --scale worker_s="${NUMWORKER}"
docker container start benchmark
docker cp benchmark:/doc ./