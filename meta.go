package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/spf13/cast"
	"time"
)

func metadataTest() {
	fmt.Printf("metadata test\n")

	config := sarama.NewConfig()
	config.Version = sarama.V2_4_0_0

	client, err := sarama.NewClient([]string{"localhost:9092", "localhost:9093", "localhost:9094"}, config)
	if err != nil {
		fmt.Printf("metadata_test try create client err :%s\n", err.Error())
		return
	}

	defer client.Close()

	// get topic set
	topics, err := client.Topics()
	if err != nil {
		fmt.Printf("try get topics err %s\n", err.Error())
		return
	}

	fmt.Printf("topics(%d):\n", len(topics))

	for _, topic := range topics {
		fmt.Println(topic)
		if topic == "__consumer_offsets" {
			continue
		}
		partitions, _ := client.Partitions(topic)
		for _, partition := range partitions {
			fmt.Println("partition ", partition)
			time := time.Now()
			offset, _ := client.GetOffset(topic, partition, cast.ToInt64(time)-1000000)
			fmt.Println("offset", offset)
		}
	}
	client.RefreshCoordinator("t user-log-to-es")
	co,_ := client.Coordinator("user-log-to-es")
	fmt.Println("codinator is ",co.Addr())

	// get broker set
	brokers := client.Brokers()
	fmt.Printf("broker set(%d):\n", len(brokers))
	for _, broker := range brokers {
		fmt.Printf("%s\n", broker.Addr())
	}
}
