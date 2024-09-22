package rabbitmq

import (
	"fmt"

	"github.com/skinkvi/crpt/internal/config"
	"github.com/skinkvi/crpt/pkg/util/logger"
	"github.com/streadway/amqp"
)

func InitRabbitMQ() (*amqp.Connection, error) {
	rabbitmqConfig := config.GetConfig().Rabbitmq

	connStr := fmt.Sprintf("amqp://%s:%s@%s:%d/", rabbitmqConfig.User, rabbitmqConfig.Password, rabbitmqConfig.Host, rabbitmqConfig.Port)

	conn, err := amqp.Dial(connStr)
	if err != nil {
		logger.GetLogger().Error(err.Error())
		return nil, err
	}

	return conn, nil
}
