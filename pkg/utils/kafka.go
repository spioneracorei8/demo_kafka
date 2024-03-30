package utils

import (
	"context"
	"log"

	"githib.com/spioneracorei8/hello_k_msg_q/config"
	"github.com/segmentio/kafka-go"
)

func KafkaConn(cfg config.KafkaConfig) *kafka.Conn {
	conn, err := kafka.DialLeader(context.Background(), "tcp", cfg.Url, cfg.Topic, 0)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	return conn
}

func IsTopicAlreadyExists(conn *kafka.Conn, topic string) bool {
	partotions, err := conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	for _, p := range partotions {
		if p.Topic == topic {
			return true
		}
	}
	return false
}
