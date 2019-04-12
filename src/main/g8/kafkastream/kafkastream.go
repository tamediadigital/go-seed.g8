package kafkastream

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

// ArticleWithUserID contains the json structure in the input kafka stream
type ArticleWithUserID struct {
	Timestamp float64 `json:"ts"`
	Request   struct {
		ArticleID string `json:"article_id"`
		Target    string `json:"target"`
	} `json:"request"`
	RequestHeaders struct {
		UserID string `json:"tda-uid"`
	} `json:"requestHeaders"`
}

// CreateKafkaConsumer creates a kafka consumer
func CreateKafkaConsumer(kafkaBrokers []string, group string, topic string) *cluster.Consumer {
	log.Println("Starting Consumer!")
	consumerConfig := cluster.NewConfig()
	consumerConfig.Consumer.Return.Errors = true
	consumerConfig.Consumer.Offsets.Initial = sarama.OffsetOldest
	consumer, err := cluster.NewConsumer(
		kafkaBrokers,
		group,
		[]string{topic},
		consumerConfig)
	if err != nil {
		log.Fatal("Couldn't configure consumer! ", err)
	}

	return consumer
}

func ProcessMessage(msg *sarama.ConsumerMessage) (output *ArticleWithUserID, err error) {
	var decodedJson ArticleWithUserID
	// Parse json
	if err := json.Unmarshal(msg.Value, &decodedJson); err != nil {
		fmt.Println(err)
		return nil, err
	}
	if decodedJson.RequestHeaders.UserID == "" {
		return nil, nil
	}
	output = &decodedJson
	return
}
