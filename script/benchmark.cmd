cd ../docker

   Rem NUMWORKER
   set /p NUMWORKER=""

   if (%NUMWORKER% == 0) (
      set /p NUMWORKER=""
   )

   Rem NUMCLUSTER
   set /p NUMCLUSTER=""

   if (%NUMCLUSTER% == 0) (
      set /p NUMCLUSTER=""
   )

   if (%NUMCLUSTER% gtr %NUMPOINT%) (
    set /p NUMCLUSTER=""
   )

   Rem Write file
   (
   echo NUMWORKER=%NUMWORKER%
   echo NUMCLUSTER=%NUMCLUSTER%
   )> ../.env


echo "NUMWORKER="%NUMWORKER%
Rem Docker
docker-compose --profile benchmark build
timeout 10
docker-compose --profile app up benchmark_s --scale worker_s=%NUMWORKER%
docker container start master
docker cp master:/doc ./