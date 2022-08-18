There are two common ways Microservices communicate with each other: Synchronous and Asynchronous. In Synchronous
communication, the caller waits for a response before sending the next message, it operates as a REST protocol on top of
HTTP. On the contrary, in Asynchronous communication, the messages are sent without waiting for a response. This is
suited for distributed systems and usually requires a message broker to manage the messages.

Choosing the Right Message Broker When choosing a broker for executing your asynchronous operations, you should consider
a few things:

Broker Scale — The number of messages sent per second in the system. Data Persistence — The ability to recover messages.
Consumer Capability — Whether the broker is capable of managing one-to-one and/or one-to-many consumers.

What is RabbitMQ(AMQP means Advanced Message Queuing Protocols)? RabbitMQ was released in 2007 and is one of the first
common message brokers to be created. It’s an open-source that delivers messages through both point-to-point and pub-sub
methods by implementing AMQP. It’s designed to support complex routing logic.

Scale: based on configuration and resources, the ballpark here is around **50K msg per second**. RabbitMQ can also
process a million messages per second but requires more resources (around 30 nodes).

Persistence: both persistent and transient messages are supported. One-to-one vs one-to-many consumers: both.

What is Kafka ? Kafka was created by LinkedIn in 2011, it is unified, high throughput, low latency processing for
handling real-time data feeds. As a distributed streaming platform, Kafka replicates a publish-subscribe service. It
provides data persistence and stores streams of records that render it capable of exchanging quality messages. Kafka
uses a binary TCP-based protocol that is optimized for efficiency and relies on a “message set” abstraction that
naturally groups messages together to reduce the overhead of the network round trip.

Scale: can send up to **a million messages** per second.

Persistence: yes. One-to-one vs one-to-many consumers: only one-to-many

Redis Redis is a bit different from the other message brokers. At its core, Redis is an in-memory data store that can be
used as either a high-performance key-value store or as a message broker. Another difference is that Redis has no
persistence but rather dumps its memory into a Disk/DB. It’s also perfect for real-time data processing.

Scale: can send up to **a million messages per second**. Persistence: basically, no — it’s an in-memory datastore.
One-to-one vs one-to-many consumers: both.

Apache Kafka: Pull-based approach. RabbitMQ: Push-based approach. In pull-based systems, the brokers waits for the
consumer to ask for data (‘pull’); if a consumer is late, it can catch up later. With push-based systems, messages are
immediately pushed to any subscribed consumer.

RabbitMQ was initially designed to be a message queue. Kafka was designed to be both a message queue and a Pub-Sub
system. A message queue is a queue-like structure where a message is published and is consumed once and only once. A
Pub-Sub system, on the other hand, allows a message to be consumed multiple times by multiple consumers.

How do kafka and RabbitMQ handle message? Kafka RabbitMQ Message ordering Provides message ordering thanks to its
partitioning. Not supported. Messages are sent to topics by message key.

Message lifetime Kafka is a log, which means that it retains messages by default. RabbitMQ is a queue, so messages are
done once It is managed this by specifying a retention policy. consumed, and acknowledgement is provided.

Delivery Guarantees Retains order only inside a partition. In a partition, Doesn't guarantee atomicity, even in relation
Kafka guarantees that the whole batch of messages either to transaction involving a single queue. fails or passes.

Message priorities N/A In RabbitMQ, you can specify message priorities and consume message with high priority first.

Short-lived Messages: Redis Redis’s in-memory database is an almost perfect fit for use-cases with short-lived messages
where persistence isn’t required. Because it provides extremely fast service and in-memory capabilities. Redis is the
perfect candidate for short retention messages where persistence isn’t so important, and you can tolerate some loss.
With the release of Redis streams in 5.0, it’s also a candidate for one-to-many use cases, which was definitely needed
due to limitations and old pub-sub capabilities.

Large Amounts of Data: Kafka Kafka is a high throughput distributed queue that’s built for storing a large amount of
data for long periods of time. Kafka is ideal for one to many use cases where persistence is required.

Complex Routing: RabbitMQ RabbitMQ is an older, yet mature broker with a lot of features and capabilities that support
complex routing. It will even support complex routing communication when the required rate is not high (more than a few
tens of thousands msg/sec). RabbitMQ can also process a million messages per second but requires more resources (around
30 nodes).

let’s dissect the components of Kafka.

1. Kafka broker and cluster
2. Producer
3. Consumer
4. Topics
5. Partition
6. Offset
7. Consumer group
8. Replica
9. Zookeeper
10. Long


1. Kafka broker and cluster Kafka is nothing but a server that manages the publishing and consumption of data. A Kafka
   server is known as a broker. A collection of brokers that maintain the same group of topics are known as a Kafka
   cluster. A cluster image is following: https://miro.medium.com/max/1372/1*ZdRc1OfbhUx6JNsSazc7KA.png


2. Producer A server that publishes data to a Kafka broker is known as a Producer. The Producer image is
   following: https://miro.medium.com/max/1372/1*Afs8Gb8PR-jJ4qwHkjoasw.png
   The Producer specifies the topic and the partition of a message before publishing.

3. Consumer A consumer is a server that subscribes and consumes data from a Kafka topic. The consumer image is
   following: https://miro.medium.com/max/1372/1*2CmxQ5fKwH3195ltMhq3nQ.png

4. Topic A Kafka broker maintains different types of events, each event is a massive stream of data. A topic is an
   append-only log that’s stored on a Kafka broker, it is simply a type of event or a stream of data. When publishing to
   Kafka, the publisher specifies the topic in which the message should be published. Appending a message to a topic is
   akin to appending data to a queue.

5. Partition Instead of storing all the data in an append-only log, a topic can be split into multiple partitions. Each
   partition stores a portion of the data for a specific topic. The topic is sharded based on partitions. This is akin
   to database sharding. A partition of the same topic can be stored either on the same or a different Kafka broker.

6. Offset An offset is a unique index of a message in a partition. An offset image is
   following https://miro.medium.com/max/1372/1*yNK9zbt4C_GCN3sHrlrY-w.png
   When Kafka pushes data to the consumer, it increases and keeps track of the current offset. Current offset and
   committed offset are two types of offset.

7. Consumer Group A consumer group consists of a group of consumers that consume the same topic. Each partition can only
   be consumed by one and only one consumer from the same group. Different groups can consume from the same topic
   concurrently with a different offset. Consumer groups know nothing about each other and consume data using a separate
   offset.

   If there’s only one consumer in a group, the consumer will be responsible for the consumption of all available
   partitions. When a new consumer joins the group, for example, a new server instance is added, Kafka will perform
   rebalancing and assigns a portion of the partitions to the new consumer. Kafka uses its own rebalance strategy for
   the partition reassignment.

8. Replica The single point of failure is the nightmare of every distributed system. Replicas are created for each
   partition and are stored on different Kafka brokers. A leader is elected for each partition to serve both publishers
   and consumers. The replicas constantly sync data from the leader. When the leader goes down, the Zookeeper joins in
   to help with the leader's election.

9. Zookeeper Zookeeper is a service synchronisation system that stores metadata and coordinates the distributed system
   in Kafka. It’s mainly involved in the following:

   Leader election — Ensure that there’s a leader for each partition Cluster membership — Keep track of all functional
   brokers in a cluster Topic configuration — Keep track of all available topics, partitions, and their replicas Access
   control list — Keep track of the number of consumers in each group and their access right Quotas — Keep track of the
   amount of data each client can read and write

10. Long Polling RabbitMQ adopts the push model. The broker maintains a persistent TCP connection with the consumers and
    pushes data to them should there be available data. A push model, however, can potentially overwhelm the consumers.
    If the brokers push data faster than the consumers can process them, the consumers might fall behind. RabbitMQ does
    have a solution for that.

    Kafka utilises a pull model, named long polling. The consumers pull data periodically from the broker. The consumers
    can pull data only when they are ready. However, if there’s no data on the partition, periodic polling from the
    consumers might result in resource wasting. Kafka does not return an empty response if there’s no data on the
    partition. Instead, the broker holds the connection and waits for data to come in before returning it to the
    consumers. That is called the "long polling". This alleviates the frequent polling from consumers when there’s no
    data on the partition and prevent the wasting of resources.

The Idempotent Consumer Pattern

An Idempotent Consumer pattern uses a Kafka consumer that can consume the same message any number of times, but only
process it once. To implement the Idempotent Consumer pattern the recommended approach is to add a table to the database
to track processed messages. Each message needs to have a unique messageId assigned by the producing service, either
within the payload, or as a Kafka message header. When a new message is consumed the table is checked for the existence
of this message Id. If exists, then the message is a duplicate. The consumer updates its offsets to effectively mark the
message as consumed to ensure it is not redelivered, no further action takes place. If the message Id is not exist in
the table then a database transaction is started and the message Id is inserted. The message is then processed
performing the required business logic. Upon completion the transaction is committed.

Possible issue when using the idempotent consumer pattern.

In the case where two duplicate events are consumed and are being processed in parallel, then both are able to start a
transaction, write (uncommitted) the message Id to the database, and perform their message processing. Only when the
second thread comes to commit the transaction would the constraint violation be thrown and the transaction be rolled
back.

Using lock and flush to fix the above issue, the first transaction lock the row or the table, after saving, do flush.

As the consumer offsets are only updated after the database transaction commits, then if the consumer dies after writing
the outbox event but before the database transaction commits, then no outbound message is sent, and the message will be
redelivered.

If the consumer dies after the database transaction commits but before the consumer offsets are written, then the
message is redelivered。

How to run and check kafka on mac?

1. % zookeeper-server-start /opt/homebrew/etc/kafka/zookeeper.properties

2. % kafka-server-start /opt/homebrew/etc/kafka/server.properties

3. % kafka-console-producer --broker-list localhost:9092 --topic ramsey

4. % kafka-console-consumer --bootstrap-server localhost:9092 --topic ramsey --from-beginning