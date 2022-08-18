ZOOKEEPER_SERVER_ID identifies each zookeper differently. KAFKA_BROKER_ID identifies each kafka broker differently.
ZOOKEEPER_SERVERS links of all zookeeper, so that they can communicate within themselves. KAFKA_ZOOKEEPER_CONNECT links
of all zookeeper

More details please click the following link: https://www.confluent.io/blog/kafka-listeners-explained/

KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: LISTENER_INTERNAL:PLAINTEXT,LISTENER_EXTERNAL:PLAINTEXT

Map our custom listener with the kafka’s security protocol either internally or externally. We will be using PLAINTEXT
security protocol. In the docker-compose.yml, LISTENER_INTERNAL and LISTENER_EXTERNAL are custom protocols, map it to
PLAINTEXT.

KAFKA_LISTENERS are mapped with KAFKA_ADVERTISED_LISTENERS sequentially.

So LISTENER_INTERNAL://kafka-1:19092 will be mapped with LISTENER_INTERNAL://kafka-1:19092 And LISTENER_EXTERNAL:
//kafka-1:19091 will be mapped with LISTENER_EXTERNAL://localhost:19091

KAFKA_INTER_BROKER_LISTENER_NAME: LISTENER_INTERNAL

Need a specific kafka which among 2 listeners to use for internal communication.

In segmentio/consumer.go, using a DialLeader initializer to initialise the consumer, then read deadline and batch size.
At the end, run a “for” loop and listen for messages and close the connection.

In segmentio/producer.go, using a DialLeader initialiser to initialise the producer, then use conn.WriteMessages, At the
end, close connection.

In segmentio pkg, the WriteMessages is not friendly as confluent and sarama.

In confluent/consumer.go, use port 19091 which was used as external communication, here is using it outside docker
environment. If in any case you want to use it within docker environment, you can use kafka-1:19092 as host. Subscribe
to a topic, here it’s myTopic. Using the “for” loop continuously listens for messages.

In confluent/producer.go, it defines a go routine function to get the delivery report of the message sent go func(). It
produces multiple messages to be consumed by consumer.

% brew install kcat

After install has been finished, run commands following one by one.

In terminal 0:

% kcat -C -b localhost:19092,localhost:29092,localhost:39092 -t foo -p 0 % Reached end of topic foo [0] at offset 0
publish to partition 0 % Reached end of topic foo [0] at offset 1

In terminal 1:
% echo 'publish to partition 0' | kcat -P -b localhost:19092,localhost:29092,localhost:39092 -t foo -p 0 % echo 'publish
to partition 1' | kcat -P -b localhost:19092,localhost:29092,localhost:39092 -t foo -p 1 % echo 'publish to partition 1'
| kcat -P -b localhost:19092,localhost:29092,localhost:39092 -t foo -p 1 % echo 'publish to partition 1' | kcat -P -b
localhost:19092,localhost:29092,localhost:39092 -t foo -p 1 % echo 'publish to partition 1' | kcat -P -b localhost:
19092,localhost:29092,localhost:39092 -t foo -p 2

In terminal 2:
% kcat -C -b localhost:19092,localhost:29092,localhost:39092 -t foo -p 1 publish to partition 1 publish to partition 1
publish to partition 1 % Reached end of topic foo [1] at offset 3

In terminal 3:
% kcat -C -b localhost:19092,localhost:29092,localhost:39092 -t foo -p 2 % Reached end of topic foo [2] at offset 0
publish to partition 1 % Reached end of topic foo [2] at offset 1