# Platform
XRPL Deep search platform is a data processing pipeline to store XRPL transactions in a distributed search and analytics platform. The platform enables simple data retrieval, aggregate information and discovering trends and patterns in XRPL. 

### What is Deep search?
XRP Ledger exploration tools and APIs available today, such as rippled, clio and various explorer APIs provide access to ledger data based on object's primary key such as ledger index, transaction hash, account address, NFT id, object ids etc. This project aims to provide deeper search capability such as filtering transactions by source/destination tags, range query over payment amounts, aggregate volumes and much more. This is enabled by indexing all properties of the transactions in an analytics engine.

### Requirements

1. [Apache Kafka](https://kafka.apache.org)
2. [rippled](https://xrpl.org/install-rippled.html)
3. Access to full history rippled (if backfilling older ledgers)
4. [Elasticsearch](https://www.elastic.co/downloads/elasticsearch)

### Architecture

![Search platform architecture](https://github.com/xrpscan/platform/blob/main/assets/xrpscan-platform.png?raw=true)

### Installation

1. This project is known to run on Linux and macOS. This README lists out steps to
run the service on CentOS.

2. Install Elasticsearch via this guide: https://www.elastic.co/guide/en/elasticsearch/reference/current/rpm.html

```
dnf install --enablerepo=elasticsearch elasticsearch
systemctl daemon-reload
systemctl enable elasticsearch.service
systemctl start elasticsearch.service
```
Elasticsearch default installer would print the instance's defauly password. This
password must be noted as it would be required in the later steps (`ELASTICSEARCH_PASSWORD`).

3. Install Docker via this guide: https://docs.docker.com/engine/install/centos/

4. Configure Docker to run as non-root

```
usermod -aG docker non-root-user
systemctl restart docker
```

5. Install Zookeeper and Kafka

```
docker compose up -d
```

6. Install Go via this guide: https://go.dev/doc/install

7. Build deep search platform

```
dnf install make
git clone git@github.com:xrpscan/platform.git
cd platform
make
```

8. Create environment file and update settings within

```
cp .env.example .env
```

9. Copy fingerprint output from the following command into .env file. `ELASTICSEARCH_FINGERPRINT="xxxxxx"`

```
openssl x509 -fingerprint -sha256 -noout -in /etc/elasticsearch/certs/http_ca.crt | sed s/://g
```

10. Create Kafka topics

```
docker exec kafka-broker1 kafka-topics --bootstrap-server kafka-broker1:9092 --create --if-not-exists --topic xrpl-platform-ledgers
docker exec kafka-broker1 kafka-topics --bootstrap-server kafka-broker1:9092 --create --if-not-exists --topic xrpl-platform-transactions
docker exec kafka-broker1 kafka-topics --bootstrap-server kafka-broker1:9092 --create --if-not-exists --topic xrpl-platform-validations
docker exec kafka-broker1 kafka-topics --bootstrap-server kafka-broker1:9092 --create --if-not-exists --topic xrpl-platform-manifests
docker exec kafka-broker1 kafka-topics --bootstrap-server kafka-broker1:9092 --create --if-not-exists --topic xrpl-platform-peerstatus
docker exec kafka-broker1 kafka-topics --bootstrap-server kafka-broker1:9092 --create --if-not-exists --topic xrpl-platform-consensus
docker exec kafka-broker1 kafka-topics --bootstrap-server kafka-broker1:9092 --create --if-not-exists --topic xrpl-platform-server
docker exec kafka-broker1 kafka-topics --bootstrap-server kafka-broker1:9092 --create --if-not-exists --topic xrpl-platform-default
docker exec kafka-broker1 kafka-topics --bootstrap-server kafka-broker1:9092 --create --if-not-exists --topic xrpl-platform-tx
```

11. Create Elasticsearch indexes

```
./bin/platform-cli init -elasticsearch -shards 8 -replicas 0
```

### Running the service

1. Index new ledgers

```
./bin/platform-server
```

2. Backfill old ledgers

```
./bin/platform-cli backfill -verbose -from 82000000 -to 82999999
```

### Monitoring the service
The services ships a command named `eps` that may be used to print Elasticsearch
index statistics. Open file `cmd/eps/eps` and update `ES_ENV_FILE` variable so
that it points to your platform `.env` file.

```
vi cmd/eps/eps  # Update ES_ENV_FILE variable
cp cmd/eps/eps /path/to/your/bin
eps
```

### Querying data
Deep search platform will provide easy to use APIs for querying XRPL transaction
data in a future release. For now, data can be queried by connecting to Elasticsearch
directly

```
source .env
curl -k -u elastic:$ELASTICSEARCH_PASSWORD \
-H 'Content-type: application/json' \
-XPOST 'https://localhost:9200/platform.transactions/_search' \
-d '{"query":{"term":{"ctid":"C511CC0400850000" }}}' | \
jq .hits
```

### References
[Ledger Stream - xrpl.org](https://xrpl.org/subscribe.html#ledger-stream)

### Known issues

- [xrpl-go tries to read from websocket even after its connection is closed](https://github.com/xrpscan/platform/issues/36)

### Reporting bugs
Please create a new issue in [Platform issue tracker](https://github.com/xrpscan/platform/issues)

### EOF