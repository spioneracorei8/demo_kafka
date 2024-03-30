package main

import (
	"fmt"
	"log"

	"githib.com/spioneracorei8/hello_k_msg_q/config"
	"githib.com/spioneracorei8/hello_k_msg_q/pkg/utils"
)

func main() {
	cfg := config.KafkaConfig{
		Url:   "localhost:9092",
		Topic: "msg-q",
	}
	conn := utils.KafkaConn(cfg)

	for {
		msg, err := conn.ReadMessage(10e3)
		if err != nil {
			break
		}
		fmt.Println(string(msg.Value))
	}
	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection: ", err.Error())
	}
}
