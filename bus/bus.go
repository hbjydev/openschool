package bus

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

type Config struct {
	Schema               string
	Username             string
	Password             string
	Host                 string
	Port                 string
	Vhost                string
	ConnectionName       string
	ChannelNotifyTimeout time.Duration
	Reconnect            struct {
		Interval   time.Duration
		MaxAttempt int
	}
}

type Bus struct {
	mux                  sync.RWMutex
	config               Config
	logger               *zap.Logger
	dialConfig           amqp.Config
	connection           *amqp.Connection
	ChannelNotifyTimeout time.Duration
}

func New(config Config, logger *zap.Logger) *Bus {
	return &Bus{
		config:               config,
		logger:               logger,
		dialConfig:           amqp.Config{Properties: amqp.Table{"connection_name": config.ConnectionName}},
		ChannelNotifyTimeout: config.ChannelNotifyTimeout,
	}
}

func (b *Bus) getConnectionString() string {
	return fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s",
		b.config.Schema,
		b.config.Username,
		b.config.Password,
		b.config.Host,
		b.config.Port,
		b.config.Vhost,
	)
}

func (b *Bus) Connect() error {
	conn, err := amqp.DialConfig(
		b.getConnectionString(),
		b.dialConfig,
	)

	if err != nil {
		return err
	}

	b.connection = conn

	go b.reconnect()

	return nil
}

func (b *Bus) Channel() (*amqp.Channel, error) {
	if b.connection == nil {
		if err := b.Connect(); err != nil {
			return nil, errors.New("no open amqp connection to bus")
		}
	}

	channel, err := b.connection.Channel()
	if err != nil {
		return nil, err
	}

	return channel, nil
}

func (b *Bus) reconnect() {
WATCH:

	conErr := <-b.connection.NotifyClose(make(chan *amqp.Error))
	if conErr != nil {
		b.logger.Error("amqp connection dropped, reconnecting...")
		var err error

		for i := 1; i <= b.config.Reconnect.MaxAttempt; i++ {
			b.mux.RLock()
			b.connection, err = amqp.DialConfig(
				b.getConnectionString(),
				b.dialConfig,
			)
			b.mux.RUnlock()

			if err == nil {
				b.logger.Info("amqp reconnected")
				goto WATCH
			}

			time.Sleep(b.config.Reconnect.Interval)
		}

		b.logger.Sugar().Errorf("failed to reconnect to amqp: %v", err.Error())
	} else {
		b.logger.Info("amqp disconnected normally, not reconnecting")
	}
}
