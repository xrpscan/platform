# Platform
XRPL Deep search platform is a data processing pipeline to store XRPL transactions in a distributed search and analytics platform. The platform enables simple data retrieval, aggregate information and discovering trends and patterns in XRPL. 

### What is Deep search?
XRP Ledger exploration tools and APIs available today, such as rippled, clio and various explorer APIs provide access to ledger data based on object's primary key such as ledger index, transaction hash, account address, NFT id, object ids etc. This project aims to provide deeper search capability such as filtering transactions by source/destination tags, range query over payment amounts, aggregate volumes and much more. This is enabled by indexing all properties of the transactions in an analytics engine.

### Requirements

1. [Apache Kafka](https://kafka.apache.org)
2. [rippled](https://xrpl.org/install-rippled.html)
3. rippled with full-history (if backfilling older ledgers)
4. [Elasticsearch](https://www.elastic.co/downloads/elasticsearch)

### Architecture

![Search platform architecture](https://github.com/xrpscan/platform/blob/main/assets/xrpscan-platform.png?raw=true)

### References
[Ledger Stream - xrpl.org](https://xrpl.org/subscribe.html#ledger-stream)

### Reporting bugs
Please create a new issue in [Platform issue tracker](https://github.com/xrpscan/platform/issues/new)

### EOF