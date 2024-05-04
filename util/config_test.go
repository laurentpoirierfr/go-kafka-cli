package util_test

import (
	"go-kafka-cli/util"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	INTEGRATION       = "test"
	RABBITMQ_USERNAME = "test"
	USER_PASSWORD     = "test"
)

type Config struct {
	Port     string   `yaml:"port"`
	RabbitMQ RabbitMQ `yaml:"rabbitmq"`
}

type RabbitMQ struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Vhost    string `yaml:"vhost"`
}

func TestConfigFile(t *testing.T) {
	err := os.Setenv("RABBITMQ_USERNAME", RABBITMQ_USERNAME)
	assert.Nil(t, err)
	err = os.Setenv("INTEGRATION", INTEGRATION)
	assert.Nil(t, err)
	err = os.Setenv("USER_PASSWORD", USER_PASSWORD)
	assert.Nil(t, err)

	config := Config{}
	err = util.LoadConfig(".", &config)
	assert.Nil(t, err)

	assert.Equal(t, "80", config.Port, "they should be equal")
	assert.Equal(t, 5672, config.RabbitMQ.Port, "they should be equal")

	assert.Equal(t, INTEGRATION, config.RabbitMQ.Vhost, "they should be equal")
	assert.Equal(t, RABBITMQ_USERNAME, config.RabbitMQ.Username, "they should be equal")
	assert.Equal(t, USER_PASSWORD, config.RabbitMQ.Password, "they should be equal")
}
