package main

import (
	"encoding/json"
	"fmt"
	"go-kafka-cli/internal/producer"
	segmentio "go-kafka-cli/internal/producer/kafka"
	"time"

	"github.com/google/uuid"
)

func main() {
	brokers := []string{"localhost:9092"}
	topic := "test-topic"

	producer := segmentio.NewKafkaMessageProducer(producer.Producer{
		Brokers: brokers,
		Topic:   topic,
	})

	fmt.Println("Démarrage du producteur de messages Kafka...")
	if err := producer.Start(); err != nil {
		fmt.Printf("Erreur lors du démarrage du producteur: %v\n", err)
		return
	}

	headers := map[string]string{"key1": "value1", "key2": "value2"}

	u := &User{
		Name:      "Sammy the Shark",
		Password:  "fisharegreat",
		CreatedAt: time.Now(),
	}

	for i := 1; i < 10000; i++ {
		msg := Message{
			ID:   uuid.New().String(),
			Meta: headers,
			Data: u,
		}
		jsonData, err := json.Marshal(msg)
		if err != nil {
			panic(err)
		}
		if err := producer.ProduceMessage(jsonData, headers); err != nil {
			fmt.Printf("Erreur lors de la production du message: %v\n", err)
		}
	}

	if err := producer.Close(); err != nil {
		fmt.Printf("Erreur lors de la fermeture du producteur: %v\n", err)
		return
	}
}

type Message struct {
	ID   string            `json:"id"`
	Meta map[string]string `json:"meta"`
	Data interface{}       `json:"data"`
}

type User struct {
	Name          string    `json:"name"`
	Password      string    `json:"password"`
	PreferredFish []string  `json:"preferredFish"`
	CreatedAt     time.Time `json:"createdAt"`
}
