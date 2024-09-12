package config

import (
	"database/sql"
	"fmt"

	internalEvents "github.com/renanmav/GoExpert-CleanArch/internal/events"
	"github.com/renanmav/GoExpert-CleanArch/internal/events/handlers"
	"github.com/renanmav/GoExpert-Events/pkg/events"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

type conf struct {
	DBDriver               string `mapstructure:"DB_DRIVER"`
	DBHost                 string `mapstructure:"DB_HOST"`
	DBPort                 string `mapstructure:"DB_PORT"`
	DBUser                 string `mapstructure:"DB_USER"`
	DBPassword             string `mapstructure:"DB_PASSWORD"`
	DBName                 string `mapstructure:"DB_NAME"`
	DBDataSourceName       string
	DB                     *sql.DB
	WebServerPort          string `mapstructure:"WEB_SERVER_PORT"`
	GrpcServerPort         string `mapstructure:"GRPC_SERVER_PORT"`
	GraphqlServerPort      string `mapstructure:"GRAPHQL_SERVER_PORT"`
	RabbitMQUser           string `mapstructure:"RABBITMQ_USER"`
	RabbitMQPassword       string `mapstructure:"RABBITMQ_PASSWORD"`
	RabbitMQHost           string `mapstructure:"RABBITMQ_HOST"`
	RabbitMQPort           string `mapstructure:"RABBITMQ_PORT"`
	RabbitMQDataSourceName string
	RabbitMQChannel        *amqp.Channel
	RabbitMQQueueName      string `mapstructure:"RABBITMQ_QUEUE_NAME"`
	RabbitMQExchange       string `mapstructure:"RABBITMQ_EXCHANGE"`
	RabbitMQRoutingKey     string `mapstructure:"RABBITMQ_ROUTING_KEY"`
	EventDispatcher        events.EventDispatcherInterface
	OrderCreatedEvent      events.EventInterface
}

func LoadConfig(path string) (cfg *conf) {
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	cfg.LoadDatabase()
	cfg.LoadRabbitMQ()
	cfg.LoadEventDispatcher()

	return cfg
}

func (c *conf) Close() {
	if c.DB != nil {
		c.DB.Close()
	}
	if c.RabbitMQChannel != nil {
		c.RabbitMQChannel.Close()
	}
}

func (c *conf) LoadDatabase() {
	c.DBDataSourceName = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
	db, err := sql.Open(c.DBDriver, c.DBDataSourceName)
	if err != nil {
		panic(err)
	}

	c.DB = db
}

func (c *conf) LoadRabbitMQ() {
	c.RabbitMQDataSourceName = fmt.Sprintf("amqp://%s:%s@%s:%s/", c.RabbitMQUser, c.RabbitMQPassword, c.RabbitMQHost, c.RabbitMQPort)
	conn, err := amqp.Dial(c.RabbitMQDataSourceName)
	if err != nil {
		panic(err)
	}

	c.RabbitMQChannel, err = conn.Channel()
	if err != nil {
		panic(err)
	}

	// Programmatically create and bind the queue
	c.RabbitMQChannel.QueueDeclare(
		c.RabbitMQQueueName, // name
		true,                // durable
		false,               // delete when unused
		false,               // exclusive
		false,               // no-wait
		nil,                 // arguments
	)
	c.RabbitMQChannel.QueueBind(
		c.RabbitMQQueueName,  // queue name
		c.RabbitMQRoutingKey, // routing key
		c.RabbitMQExchange,   // exchange
		false,                // no-wait
		nil,                  // arguments
	)
}

func (c *conf) LoadEventDispatcher() {
	c.EventDispatcher = events.NewEventDispatcher()

	orderCreatedEvent := internalEvents.NewOrderCreated()
	c.OrderCreatedEvent = orderCreatedEvent

	orderCreatedHandler := handlers.NewOrderCreatedHandler(
		c.RabbitMQChannel,
		c.RabbitMQExchange,
		c.RabbitMQRoutingKey,
	)

	c.EventDispatcher.Register("OrderCreated", orderCreatedHandler)
}
