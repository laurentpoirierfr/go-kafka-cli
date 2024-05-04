package kafka

import (
	"context"
	"go-kafka-cli/internal/producer"

	"github.com/segmentio/kafka-go"
)

// KafkaMessageProducer est une implémentation de MessageProducer pour Kafka.
type KafkaMessageProducer struct {
	brokers []string
	topic   string
	writer  *kafka.Writer
}

// NewKafkaMessageProducer crée une nouvelle instance de KafkaMessageProducer.
func NewKafkaMessageProducer(p producer.Producer) *KafkaMessageProducer {
	return &KafkaMessageProducer{
		brokers: p.Brokers,
		topic:   p.Topic,
	}
}

// Start initialise le producteur de messages Kafka.
func (p *KafkaMessageProducer) Start() error {
	p.writer = kafka.NewWriter(kafka.WriterConfig{
		Brokers:  p.brokers,
		Topic:    p.topic,
		Balancer: &kafka.LeastBytes{},
		Async:    true,
	})

	return nil
}

// ProduceMessage produit un message Kafka avec des en-têtes personnalisés.
func (p *KafkaMessageProducer) ProduceMessage(message []byte, headers map[string]string) error {
	var kafkaHeaders []kafka.Header
	for key, value := range headers {
		kafkaHeaders = append(kafkaHeaders, kafka.Header{Key: key, Value: []byte(value)})
	}

	return p.writer.WriteMessages(context.Background(),
		kafka.Message{
			Value:   message,
			Headers: kafkaHeaders,
		},
	)
}

// Close arrête proprement le producteur de messages Kafka.
func (p *KafkaMessageProducer) Close() error {
	if p.writer != nil {
		return p.writer.Close()
	}
	return nil
}
