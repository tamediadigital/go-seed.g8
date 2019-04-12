package g8

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"strconv"

	"github.com/Shopify/sarama"
	"github.com/go-redis/redis"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/tamediadigital/golang-stream-processor-sample/config"
	"github.com/tamediadigital/golang-stream-processor-sample/kafkastream"
)

var (
	processedClickCount = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "processed_click_count",
			Help: "Click count",
		})
	errorCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "error_count",
			Help: "Error count",
		}, []string{"error_type"})
)

func ReportError(err error) {
	log.Println(err.Error())
	errorCount.WithLabelValues(
		reflect.TypeOf(err).String()).Inc()
}

func SetupRedis(cfg config.Config) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisHost + ":" + strconv.Itoa(cfg.RedisPort),
		Password: "",
		DB:       cfg.RedisDB,
	})

	// Test redis connection, since above doesn't actually try to connect to redis
	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatal("Redis connection failure! ", err)
	}

	return redisClient
}

func ProcessMessage(msg *sarama.ConsumerMessage, redisClient *redis.Client, cfg config.Config) error {
	userData, err := kafkastream.ProcessMessage(msg)
	if err != nil {
		return err
	}
	processedClickCount.Inc()
	key := cfg.RedisUsersKey + "/" + userData.Request.Target
	field := userData.Request.ArticleID
	return redisClient.HIncrBy(key, field, 1).Err()
}

func RunProcessor(cfg config.Config, redisClient *redis.Client) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// Setup Kafka input
	consumer := kafkastream.CreateKafkaConsumer(cfg.KafkaBrokers, cfg.KafkaConsumerGroup, cfg.KafkaTopic)
	defer consumer.Close()

	// consume messages, watch signals
	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				if err := ProcessMessage(msg, redisClient, cfg); err != nil {
					ReportError(err)
				} else {
					processedClickCount.Inc()
				}

				// mark message as processed
				consumer.MarkOffset(msg, "")
			} else {
				log.Println("Message read failed!")
			}
		case err := <-consumer.Errors():
			ReportError(err)
		case <-signals:
			return
		}
	}
}

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(processedClickCount)
	prometheus.MustRegister(errorCount)
}

func main() {
	log.Println("Starting application!")
	log.Println("Loading config!")
	cfg := config.LoadConfiguration()

	log.Println("Starting metrics exporting!")
	http.Handle("/metrics", promhttp.Handler())
	go func() { log.Fatal(http.ListenAndServe(cfg.PrometheusEndpoint, nil)) }()

	log.Println("Setting up redis!")
	redisClient := SetupRedis(cfg)

	log.Println("Starting stream processing!")
	RunProcessor(cfg, redisClient)
}
