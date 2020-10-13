package main

import (
	"github.com/yaminmhd/go-kafka-producer-protobuf/cmd"
	"github.com/yaminmhd/go-kafka-producer-protobuf/config"
)

func main() {
	config.Load()
	startApp()
}

func startApp() {
	cmd.ProduceMessage()
}
