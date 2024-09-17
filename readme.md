# Kafka with Go: Complete Documentation with Docker Setup

### Use cases
- **Messaging:**  In comparison to most messaging systems Kafka has better throughput, built-in partitioning, replication, and fault-tolerance which makes it a good solution for large scale message processing applications.
- **Website Activity Tracking:** The original use case for Kafka was to be able to rebuild a user activity tracking pipeline as a set of real-time publish-subscribe feeds.
- **Metrics:** Kafka is often used for operational monitoring data. This involves aggregating statistics from distributed applications to produce centralized feeds of operational data.
- **Log Aggregation:** Log aggregation typically collects physical log files off servers and puts them in a central place (a file server or HDFS perhaps) for processing. Kafka abstracts away the details of files and gives a cleaner abstraction of log or event data as a stream of messages.
- **Stream Processing:** Many users of Kafka process data in processing pipelines consisting of multiple stages, where raw input data is consumed from Kafka topics and then aggregated, enriched, or otherwise transformed into new topics for further consumption or follow-up processing.
- **Event Sourcing:** Event sourcing is a style of application design where state changes are logged as a time-ordered sequence of records. Kafka's support for very large stored log data makes it an excellent backend for an application built in this style.
- **Commit Log:** Kafka can serve as a kind of external commit-log for a distributed system.

### Table of contents
1. Docker setup for kafka
2. Go kafka producer
3. Go kafka consumer
4. Advanced kafka use cases
   - Adding headers to kafka messages
   - Custom partitioning
   - Handling offsets in consumer

### 1. Docker Setup for Kafka
#### 1.1. Docker compise setup
```yaml
version: '2'
services:
  zookeeper:
    image: confluentinc/cp-zookeeper:latest
    environment:
      ZOOKEEPER_CLIENT_PORT: 2181
      ZOOKEEPER_TICK_TIME: 2000
    ports:
      - "2181:2181"

  kafka:
    image: confluentinc/cp-kafka:latest
    depends_on:
      - zookeeper
    ports:
      - "9092:9092"
    environment:
      KAFKA_BROKER_ID: 1
      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://localhost:9092
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 1

```

#### 1.2. Start Kafka and Zookeeper
```bash
docker-compose up -d
```
You should see both the Kafka and Zookeeper containers running.

### 2. Go Kafka Producer

#### 2.1. Install the Sarama Kafka Library

```bash
go get github.com/IBM/sarama
```

#### 2.2. Implementing the Producer
[producer.go](https://github.com/shabrul2451/Kafka-with-go/blob/main/producer/producer.go)

#### 2.3. Run the Producer

```
go run producer.go
```

This will send a message to the Kafka topic <kafka-topic>.

### 3. Go Kafka Consumer

#### 3.1. Implementing the consumer

[consumer.go](https://github.com/shabrul2451/Kafka-with-go/blob/main/consumer/consumer.go)

#### 3.2. Run the consumer
```
go run consumer.go
```
this will listen for messages from the <kafka-topic>

### 4. Advanced kafka use cases

#### 4.1. Adding headers to kafka Messages

[producer-with-custom-header.go](https://github.com/shabrul2451/Kafka-with-go/blob/main/producer/producer-with-custom-header.go)

you can add as much header as you want.

#### 4.2. Custom Partitioning

By default, Kafka decides which partition to send the message to, based on the message key. However, you can explicitly set the partition in the ProducerMessage.

[producer-with-custom-partition.go](https://github.com/shabrul2451/Kafka-with-go/blob/main/producer/producer-with-custom-partition.go)
