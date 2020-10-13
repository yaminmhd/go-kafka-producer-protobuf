## About 
This application allows users to push message serialized using protobuf into kafka

- An example proto is provided in the `protos` directory. 
- To compile and generate GO codes from the proto file, use the command `make generate`. This command
 will generate a `*.pb.go` in the `generatedProtos/person` directory.
- A sample generatedProto is already provided in the `generatedProtos` directory.
- To run application locally, ensure zookeeper and kafka server is running, then run with command `make run`.
- Kafka message will be published to local kafka, you can read the message and print to console using another
library [kafka-protobuf-console-consumer](https://github.com/yogeshsr/kafka-protobuf-console-consumer).
This library is able to read protobuf encoded messages from a kafka topic and prints its decoded JSON to console.

## Usage
Generate proto files
- `make generate`

Build files and create binary output
- `make build`

Run
- `make run`