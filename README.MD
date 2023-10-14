## create topic in kafka
1. docker exec -it kafka_1 sh
2. kafka-topics --create --bootstrap-server localhost:9093 --replication-factor 1 --partitions 1 --topic sendmail --command-config /etc/kafka/configs/config.properties
