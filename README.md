# Retry Mechanism

## Producer
The producer includes a retry mechanism to handle transient errors when producing messages to the Kafka topic. This ensures that if a message fails to be sent initially, the producer will retry sending it a configurable number of times before giving up.

## Consumer
The consumer also includes a retry mechanism for sending emails via SMTP. If an email fails to send, the consumer will retry sending the email a configurable number of times before logging the failure and moving on. 
Additionally, the consumer includes a retry mechanism for committing messages to Kafka. If a commit fails, it will retry a configurable number of times before logging the failure.

# Kafka Setup Guide

## Running kafka docker
1. Navigate to the kafka directory and run the kafka docker compose:
   ```shell
    cd kafka && docker compose up
   ```
## Creating a Topic in Kafka

1. Access the Kafka container:
   ```sh
   docker exec -it kafka_1 sh
   ```
2. create topic kafka
   ```sh
   kafka-topics --create --bootstrap-server localhost:9093 --replication-factor 1 --partitions 1 --topic sendmail --command-config /etc/kafka/config/config.properties
    ```
   

# Running

##  Run producer
1. Copy the .env.example file to .env in the producer directory:
   ```shell
    cp .env.example .env
   ```
2. Update the .env file with your configuration settings.
3. Navigate to the producer directory and run the producer:
   ```shell
   cd producer && 
   go run main.go
   ```

## Running consumer
1. Copy the .env.example file to .env in the consumer directory:
   ```shell
    cp .env.example .env
   ```
2. Update the .env file with your configuration settings. You can use Mailtrap or other SMTP services.
3. Navigate to the consumer directory and run the consumer:
   ```shell
   cd consumer && 
   go run main.go
   ```