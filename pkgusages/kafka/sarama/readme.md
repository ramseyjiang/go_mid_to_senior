KAFKA_CREATE_TOPICS: "sarama_topic:2:1"
It means a topic called ‘sarama_topic’ was created with 2 partitions and 1 replica by an environment variable.

To start the containers, go to the docker-compose directory and run docker-compose up -d

config.Version = sarama.MaxVersion It means current version is sarama the latest version. In Goland, after input ".", it
will show an option list.

As our topic has two partitions, the producer is sending messages for both of them, and our consumer group (that has
only one consumer) is consuming from both partitions.

To start another consumer in the same consumer group we just need to open a new terminal. It is important to keep the
first consumer running and run go run main.go consumer. Note that when the Kafka identifies that there’s a new consumer
in this group, a coordinator is called to balance the load between all consumers in this consumer group.

It’s possible to see that your first consumer will log something like the following block:

**Before open another terminal:**
2022/08/18 16:07:54 Message topic:"sarama_topic" partition:0 offset:36 message: testing 1234 2022/08/18 16:07:55 Message
topic:"sarama_topic" partition:1 offset:32 message: testing 1234 2022/08/18 16:07:56 Message topic:"sarama_topic"
partition:0 offset:37 message: testing 1234 2022/08/18 16:07:57 Message topic:"sarama_topic" partition:1 offset:33
message: testing 1234 2022/08/18 16:07:58 Message topic:"sarama_topic" partition:1 offset:34 message: testing 1234

**After open another terminal:**
[sarama_logger]2022/08/18 16:07:59 consumergroup/sarama_consumer loop check partition number coroutine will exit,
topics [sarama_topic]
[sarama_logger]2022/08/18 16:07:59 consumer/broker/1001 closed dead subscription to sarama_topic/0
[sarama_logger]2022/08/18 16:07:59 consumer/broker/1001 closed dead subscription to sarama_topic/1
[sarama_logger]2022/08/18 16:07:59 consumergroup/session/sarama-7ad0e968-451d-409e-921d-6f9fa16ae06a/16 heartbeat loop
stopped
[sarama_logger]2022/08/18 16:07:59 consumergroup/session/sarama-7ad0e968-451d-409e-921d-6f9fa16ae06a/16 released
[sarama_logger]2022/08/18 16:07:59 client/metadata fetching metadata for [sarama_topic] from broker localhost:9093
[sarama_logger]2022/08/18 16:07:59 client/coordinator requesting coordinator for consumergroup sarama_consumer from
localhost:9093
[sarama_logger]2022/08/18 16:07:59 client/coordinator coordinator for consumergroup sarama_consumer is #1001 (localhost:
9093)
[sarama_logger]2022/08/18 16:07:59 consumer/broker/1001 accumulated 1 new subscriptions
[sarama_logger]2022/08/18 16:07:59 consumer/broker/1001 added subscription to sarama_topic/1 2022/08/18 16:08:00 Message
topic:"sarama_topic" partition:1 offset:35 message: testing 1234