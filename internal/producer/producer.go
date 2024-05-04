package producer

// MessageProducer définit une interface pour la production de messages Kafka.
type MessageProducer interface {
	Start() error
	ProduceMessage(message []byte, headers map[string]string) error
	Close() error
}

type Producer struct {
	Brokers []string
	Topic   string
}
