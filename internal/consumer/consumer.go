package consumer

// MessageConsumer définit une interface pour la consommation de messages Kafka.
type MessageConsumer interface {
	Start() error
	ConsumeMessages(handleMessage func(msg []byte, headers map[string]string) error) <-chan error
	Close() error
}

type Consumer struct {
	Brokers     []string
	Topic       string
	GroupID     string
	StartOffset string
}

const (
	Latest   string = "Latest"
	Earliest        = "Earliest"
)
