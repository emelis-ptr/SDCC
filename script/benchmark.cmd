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

Rem ALGO
CHOICE /C 123 /M "Select [1]: LLyod, [2]: standard kmeans, [3]: keans plus plus"
set CHOICE=%errorlevel%

if %CHOICE%==1 (
 set ALGO=llyod
)
if %CHOICE%==2 (
 set ALGO=standardKMeans
)
if %CHOICE%==3 (
 set ALGO=kmeansPlusPlus
)

   Rem Write file
   (
   echo NUMWORKER=%NUMWORKER%
   echo NUMCLUSTER=%NUMCLUSTER%
   echo ALGO=%ALGO%
   )> ../.env


echo "NUMWORKER="%NUMWORKER%
Rem Docker
docker-compose --profile benchmark build
timeout 10
docker-compose --profile app up benchmark_s --scale worker_s=%NUMWORKER%
docker container start benchmark
docker cp benchmark:/doc ./