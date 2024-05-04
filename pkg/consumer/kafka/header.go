package kafka

import "github.com/segmentio/kafka-go"

// kafkaHeadersCarrier est un adaptateur permettant d'utiliser les en-têtes Kafka comme transport pour les en-têtes OpenTelemetry.
type KafkaHeadersCarrier struct {
	headers []kafka.Header
}

func NewKafkaHeadersCarrier(headers []kafka.Header) *KafkaHeadersCarrier {
	return &KafkaHeadersCarrier{headers: headers}
}

// Get retourne la valeur de l'en-tête avec la clé donnée.
func (c KafkaHeadersCarrier) Get(key string) string {
	for _, h := range c.headers {
		if h.Key == key {
			return string(h.Value)
		}
	}
	return ""
}

// Set définit la valeur de l'en-tête avec la clé donnée.
func (c KafkaHeadersCarrier) Set(key, value string) {
	// Ne fait rien, car nous ne récupérons que les valeurs des en-têtes
}

// Keys retourne les clés de tous les en-têtes.
func (c KafkaHeadersCarrier) Keys() []string {
	keys := make([]string, len(c.headers))
	for i, h := range c.headers {
		keys[i] = string(h.Key)
	}
	return keys
}

// HeadersToMap convertit un tableau de kafka.Header en une carte map[string]string.
func (c KafkaHeadersCarrier) ToMap() map[string]string {
	headerMap := make(map[string]string)
	for _, header := range c.headers {
		headerMap[string(header.Key)] = string(header.Value)
	}
	return headerMap
}
