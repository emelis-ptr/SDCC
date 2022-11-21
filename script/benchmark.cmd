cd ../docker

CHOICE /C yn /M "Benchmark after execution main [y] - Benchmark with no execution main[n]"
set BENCHMARK=%errorlevel%

if %BENCHMARK% == 2 goto :cond
goto :skip
:cond
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
:skip

echo "NUMWORKER="%NUMWORKER%
Rem Docker
docker-compose --profile benchmark build
timeout 10
docker-compose --profile app up benchmark_s --scale worker_s=%NUMWORKER%