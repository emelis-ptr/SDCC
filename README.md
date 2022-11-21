# Sistemi Distribuiti e Cloud Computing
## Algoritmo di clustering k-means in stile MapReduce e in Go

Lo scopo del progetto è realizzare nel linguaggio di programmazione Go un’applicazione distribuita che implementi
l’algoritmo di clustering k-means in versione distribuita secondo il paradigma di computazione
MapReduce.

Sono stati implementati tre diversi algoritimi kmeans: **_llyod, kmeans standard e kmeans++_**.
### Running
Sul file di script run, viene avviata l'esecuzione dei container dell'applicazione implementata: 
```
docker-compose up -d master_s --scale worker_s=%NUMWORKER% 
```

Sul file di script benchmark, viene avviata l'esecuzione del benchmark (il main in questo caso agisce come
master ma esegue i test con un numero diverso di punti, mapper e reducer):
```
docker compose --profile app up benchmark_s --scale worker_s=%NUMWORKER%
```

#### Locale
Per eseguire l'applicazione in locale è necessario aver installato:
- Docker Compose per Windows

All'interno della cartella script, esiste un file necessario per l'esecuzione. 
- Per avviare: `run.cmd`.
- Per terminare ed eliminare i container: `stop.cmd`

#### EC2
Per eseguire l'applicazione su un'istanza EC2, è necessario:
- AWS Cli
- Git bash

Dopo aver avviato un'istanza EC2 e aver salvato la chiave in locale, bisogna:
- Aprire un client SSH, nel nostro caso git bash 
- Individuare il file della chiave privata. La chiave utilizzata per avviare questa istanza è key-sdcc.pem 

Eseguire:
``` 
#il comando per far sì che la chiave non sia visualizzabile pubblicamente. 
chmod 400 key-sdcc.pem

#Connettersi all'istanza: 
ssh -i key-sdcc.pem ec2-user@<IndirizzoIP dell'istanza>
```

Una volta connessi all'istanza:
```
#installare docker
sudo yum install -y docker

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

#rieseguire il comando
ssh -i key-sdcc.pem ec2-user@<IndirizzoIP dell'istanza>
```

E' possibile eseguire l'applicazione attraverso lo script con il comando:
 - Per avviare: `sh run.sh`
 - Per terminare ed eliminare i container: ` sh stop.sh`



