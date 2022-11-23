# Sistemi Distribuiti e Cloud Computing
## Algoritmo di clustering k-means in stile MapReduce e in Go

Lo scopo del progetto è realizzare nel linguaggio di programmazione Go un’applicazione distribuita che implementi
l’algoritmo di clustering k-means in versione distribuita secondo il paradigma di computazione
MapReduce.

Sono stati implementati tre diversi algoritimi kmeans: **_llyod, kmeans standard e kmeans++_**.
### Running
Sul file di script [run](script/run.cmd), viene avviata l'esecuzione dei container dell'applicazione implementata: 
```
docker-compose up -d master_s --scale worker_s=%NUMWORKER% 
```

Sul file di script [benchmark](script/benchmark.cmd), viene avviata l'esecuzione del benchmark (il main in questo caso agisce come
master ma esegue i test con un numero diverso di punti, mapper e reducer):
```
docker compose --profile app up benchmark_s --scale worker_s=%NUMWORKER%
```

#### Locale
Per eseguire l'applicazione in locale è necessario aver installato:
- Docker Compose per Windows

All'interno della cartella script, esiste un file necessario per l'esecuzione. 
- Per avviare: [`run.cmd`](script/run.cmd)
- Per terminare ed eliminare i container: [`stop.cmd`](script/stop.cmd)

#### EC2
Per eseguire l'applicazione su un'istanza EC2, è necessario:
- AWS Cli
- Git bash

Su aws, bisogna:
- creare la chiave privata con nome "key-sdcc" e salvarla all'interno della cartella script
- creare una security group con regole di entrata e uscita:
  - Regola 1: SSH con porta 22 e destinazione: 0.0.0.0/0
  - Regola 2: HTTPS con porta 443 e destinazione: 0.0.0.0/0

Per creare un'istanza su aws: [`create_ec2.sh`](script/create_ec2.sh). Bisogna specificare
all'interno del file: AMI e security groud ID.

Eseguire:
``` 
#inserire AWS Acces Key ID, AWS Secret Access Key e region name: eu-cental-1
aws configure

#il comando per far sì che la chiave non sia visualizzabile pubblicamente. 
chmod 400 key-sdcc.pem
```

Connettersi all'istanza: 
- `ssh -i key-sdcc.pem ec2-user@<IndirizzoIP dell'istanza>`

Una volta connessi all'istanza:
```
#installare docker
sudo yum install -y docker
sudo yum install -y git

#eseguire docker
sudo service docker start

#installare plugin docker-compose manualmente:
DOCKER_CONFIG=${DOCKER_CONFIG:-$HOME/.docker}
mkdir -p $DOCKER_CONFIG/cli-plugins
curl -SL https://github.com/docker/compose/releases/download/v2.12.2/docker-compose-linux-x86_64 -o $DOCKER_CONFIG/cli-plugins/docker-compose
chmod +x $DOCKER_CONFIG/cli-plugins/docker-compose

#includere ec2-user nelle configurazioni di docker in modo tale da eseguire i comandi senza sudo: 
sudo usermod -a -G docker ec2-user

#bisogna eseguire comando per far si che le configurazioni vengano applicate
exit 
```

Rieseguire il comando:
 - `ssh -i key-sdcc.pem ec2-user@<IndirizzoIP dell'istanza>`

```
# clonare repository
git clone https://github.com/emelis-ptr/SDCC-project.git

# cartella contenente script
cd SDCC-project/script

# eseguire script
sh run.sh
```

E' possibile eseguire l'applicazione attraverso lo script con il comando:
 - Per avviare: [`sh run.sh`](script/run.sh)
 - Per terminare ed eliminare i container: [`sh stop.sh`](script/stop.sh)

Per terminare l'istanza su aws, eseguire il comando:
`aws ec2 terminate-instances --instance-ids [Instanza ID]`

#### Benchmark
Per testare le prestazioni dell'applicazione, avviare lo script:
- [`sh benchmark.sh`](script/benchmark.sh), o
- [`benchmark.cmd`](script/benchmark.cmd)

