# Project Name: Microservice for Toll Calculation

## Description
This project is a microservice for toll calculation written in Go. It utilizes Kafka for its operations.

## Getting Started
To start the project, follow these steps:

1. Run Kafka by following the instructions in the [Confluent Kafka Quick Start Guide](https://developer.confluent.io/quickstart/kafka-local/).

2. Run the microservices in the following order:
   - Run `make agg`
   - Run `make calculator`
   - Run `make receiver`
   - Run `make obu`

By following these steps, you can successfully run the microservice for toll calculation written in Go.
