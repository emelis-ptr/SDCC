cd ..

set /p NUMWORKER=

if (%NUMWORKER% == 0) (
   set /p NUMWORKER=
)

set /p NUMPOINT=

if (%NUMPOINT% == 0) (
   set /p NUMPOINT=
)

set /p NUMCLUSTER=

if (%NUMCLUSTER% == 0) (
   set /p NUMCLUSTER=
)

set /p NUMMAPPER=

if (%NUMMAPPER% == 0) (
   set /p NUMMAPPER=
)

set /p NUMREDUCER=

if (%NUMREDUCER% == 0) (
   set /p NUMREDUCER=
)

if (%NUMCLUSTER% gtr %NUMPOINT%) (
 set /p NUMCLUSTER=
)

CHOICE /C 123 /M "Select [1]: LLyod, [2]: standard kmeans, [3]: keans plus plus"
set choice=%errorlevel%

if %choice%==1 (
 set ALGO=llyod
)
if %choice%==2 (
 set ALGO=standardKMeans
)
if %choice%==3 (
 set ALGO=kmeansAlgo
)

echo NUMWORKER=%NUMWORKER% > .env
echo NUMPOINT=%NUMPOINT%>> .env
echo NUMCLUSTER=%NUMCLUSTER%>> .env
echo ALGO=%ALGO%>> .env
echo NUMMAPPER=%NUMMAPPER%>> .env
echo NUMREDUCER=%NUMREDUCER%>> .env

docker-compose build

timeout 10

docker-compose up --scale worker_s=%NUMWORKER%