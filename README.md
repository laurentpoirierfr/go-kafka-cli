# go-kafka-cli


## Producer

```golang
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
```

## Consumer

```golang
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
```
