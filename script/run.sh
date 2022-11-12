#!/bin/bash
cd ..

NUMWORKER=3
NUMPOINT=100
NUMCLUSTER=30
ALGO=""

echo ${NUMWORKER}
echo ${NUMPOINT}
echo ${NUMCLUSTER%}

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

echo NUMWORKER=${NUMWORKER} > .env
echo NUMPOINT=${NUMPOINT}>> .env
echo NUMCLUSTER=${NUMCLUSTER}>> .env
echo ALGO=${ALGO}>> .env

docker compose build

sleep 10

docker compose up --scale worker_s=${NUMWORKER}