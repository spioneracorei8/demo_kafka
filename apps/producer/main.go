package main

import (
	"log"
	"time"

	"githib.com/spioneracorei8/hello_k_msg_q/config"
	"githib.com/spioneracorei8/hello_k_msg_q/models"
	"githib.com/spioneracorei8/hello_k_msg_q/pkg/utils"
	"github.com/segmentio/kafka-go"
)

func main() {
	// Config part
	cfg := config.KafkaConfig{
		Url:   "localhost:9092",
		Topic: "msg-q",
	}
	conn := utils.KafkaConn(cfg)
	if !utils.IsTopicAlreadyExists(conn, cfg.Topic) {
		topicCfg := []kafka.TopicConfig{
			{
				Topic:             cfg.Topic,
				NumPartitions:     1,
				ReplicationFactor: 1,
			},
		}

		err := conn.CreateTopics(topicCfg...)
		if err != nil {
			panic(err.Error())
		}
	}
	data := func() []kafka.Message {
		msgs := []models.Message{
			{
				ID:          1,
				Message:     "msg-1",
				Author:      "robot-1",
				CreatedDate: time.Now(),
			},
			{
				ID:          2,
				Message:     "msg-2",
				Author:      "robot-2",
				CreatedDate: time.Now(),
			},
			{
				ID:          3,
				Message:     "msg-3",
				Author:      "robot-3",
				CreatedDate: time.Now(),
			},
		}
		messages := make([]kafka.Message, 0)
		for _, p := range msgs {
			messages = append(messages, kafka.Message{
				Value: utils.CompressToJson(p),
			})
		}
		return messages
	}()

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err := conn.WriteMessages(data...)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
