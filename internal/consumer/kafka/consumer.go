package kafka

import (
	"context"
	"fmt"
	"go-kafka-cli/internal/consumer"
	"os"
	"os/signal"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	brokers     []string
	topic       string
	groupID     string
	startOffset string
}

// NewKafkaConsumer crée une nouvelle instance de KafkaConsumer.
func NewKafkaConsumer(c consumer.Consumer) *KafkaConsumer {
	return &KafkaConsumer{
		brokers:     c.Brokers,
		topic:       c.Topic,
		groupID:     c.GroupID,
		startOffset: c.StartOffset,
	}
}

// Start initialise le consommateur de messages Kafka.
func (k *KafkaConsumer) Start() error {
	return nil // Aucune initialisation nécessaire avec go-kafka
}

// ConsumeMessages commence à consommer les messages Kafka dans une goroutine.
func (k *KafkaConsumer) ConsumeMessages(handleMessage func(msg []byte, headers map[string]string) error) <-chan error {
	errCh := make(chan error)

	go func() {
		defer close(errCh)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, os.Interrupt)

		conn, err := kafka.DialContext(ctx, "tcp", k.brokers[0])
		if err != nil {
			errCh <- err
			return
		}
		defer conn.Close()

		var startOffset int64
		if k.startOffset == consumer.Latest {
			startOffset = kafka.LastOffset
		} else {
			startOffset = kafka.FirstOffset
		}

		reader := kafka.NewReader(kafka.ReaderConfig{
			Brokers:     k.brokers,
			Topic:       k.topic,
			GroupID:     k.groupID,
			StartOffset: startOffset,
			// Logger:      kafka.LoggerFunc(log.New(log.Writer(), "[legit consumer] ", 0).Printf),
			// MinBytes:    10e3, // 10KB
			// MaxBytes:    10e6, // 10MB
			// MaxWait:     2 * time.Second,

		})
		defer reader.Close()

		for {
			select {
			case <-sigchan:
				fmt.Println("Arrêt du consommateur de messages Kafka...")
				return
			default:
				m, err := reader.ReadMessage(ctx)
				if err != nil {
					errCh <- err
					return
				}
				//fmt.Printf("Message reçu: %s\n", string(m.Value))
				headers := NewKafkaHeadersCarrier(m.Headers)
				if err := handleMessage(m.Value, headers.ToMap()); err != nil {
					errCh <- err
					return
				}
			}
		}
	}()

	return errCh
}

// Close arrête proprement le consommateur de messages Kafka.
func (k *KafkaConsumer) Close() error {
	// Aucune action nécessaire pour fermer la connexion Kafka dans go-kafka
	return nil
}
