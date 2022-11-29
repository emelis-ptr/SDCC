cd ../docker

Rem NUMWORKER
set /p NUMWORKER=""

if (%NUMWORKER% == 0) (
   set /p NUMWORKER=""
)

Rem NUMPOINT
set /p NUMPOINT=""

if (%NUMPOINT% == 0) (
   set /p NUMPOINT=""
)

Rem NUMCLUSTER
set /p NUMCLUSTER=""

if (%NUMCLUSTER% == 0) (
   set /p NUMCLUSTER=""
)

if (%NUMCLUSTER% gtr %NUMPOINT%) (
 set /p NUMCLUSTER=""
)

Rem NUMMAPPER
set /p NUMMAPPER=""

if (%NUMMAPPER% == 0) (
   set /p NUMMAPPER=""
)

Rem NUMREDUCER
set /p NUMREDUCER=""

if (%NUMREDUCER% == 0) (
   set /p NUMREDUCER=""
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
echo NUMPOINT=%NUMPOINT%
echo NUMCLUSTER=%NUMCLUSTER%
echo NUMMAPPER=%NUMMAPPER%
echo NUMREDUCER=%NUMREDUCER%
echo ALGO=%ALGO%
)> ../.env

Rem docker
docker-compose --profile app build
timeout 10
docker-compose up master_s --scale worker_s=%NUMWORKER%
docker container start master
docker cp master:/doc ./