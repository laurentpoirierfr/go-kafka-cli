package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-kafka-cli/internal/consumer"
	segmentio "go-kafka-cli/internal/consumer/kafka"
)

func main() {
	brokers := []string{"localhost:9092"}
	topic := "test-topic"
	groupID := "test-group"

	c := consumer.Consumer{
		Brokers:     brokers,
		Topic:       topic,
		GroupID:     groupID,
		StartOffset: consumer.Earliest,
	}

	cons := segmentio.NewKafkaConsumer(c)

	fmt.Println("Démarrage de la consommation de messages Kafka...")
	if err := cons.Start(); err != nil {
		fmt.Printf("Erreur lors du démarrage du consommateur: %v\n", err)
		return
	}

	errCh := cons.ConsumeMessages(HandleMessage)

	err := <-errCh
	if err != nil {
		fmt.Printf("Erreur lors de la fermeture du consommateur: %v\n", err)
		return
	}

	// // Attente d'erreurs en provenance de la goroutine de consommation
	// select {
	// case err := <-errCh:
	// 	fmt.Printf("Erreur lors de la consommation de messages: %v\n", err)
	// case <-time.After(30 * time.Second): // Attendre pendant une durée arbitraire pour l'exemple
	// 	fmt.Println("La consommation de messages s'est terminée sans erreur.")
	// }

	if err := cons.Close(); err != nil {
		fmt.Printf("Erreur lors de la fermeture du consommateur: %v\n", err)
		return
	}

	fmt.Println("La consommation de messages s'est terminée sans erreur.")
}

func HandleMessage(msg []byte, headers map[string]string) error {
	// Votre logique de traitement de message ici
	// Cette fonction retourne une erreur si le message doit être acquitté,
	// sinon, retournez nil pour passer le message à un autre topic.

	if len(msg) == 0 {
		// Exemple: Si le message est vide, nous le rejetons
		return errors.New("message vide, rejeté")
	}

	message := Message{}

	err := json.Unmarshal(msg, &message)
	if err != nil {
		fmt.Println("marshaling error: %s", err)
	}

	marshaled, err := json.MarshalIndent(message, "", "\t")
	if err != nil {
		fmt.Println("marshaling error: %s", err)
	}
	fmt.Println(string(marshaled))

	// Dans cet exemple, nous acquittons toujours le message
	fmt.Println("Acquitter le message")
	return nil
}

type Message struct {
	ID   string            `json:"id"`
	Meta map[string]string `json:"meta"`
	Data interface{}       `json:"data"`
}
