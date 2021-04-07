package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	ftK "github.com/feitianlove/golib/kafka"
	"github.com/feitianlove/logtransfers/config"
	"github.com/feitianlove/logtransfers/kafka"
	"strconv"
	"sync"
	"time"
)

//TODO golib 中初始化kafka变成 string
func main() {
	var wg sync.WaitGroup
	cfg, err := config.NewConfig("./etc/logtransfer.conf")
	if err != nil {
		panic(err)
	}
	fmt.Printf("config init success: %+v", cfg.Kafka)
	kconfig := ftK.Kafka{ServerAddr: cfg.Kafka.Address}
	produce, err := kafka.InitProduct(kconfig)
	if err != nil {
		panic(err)
	}
	consumer, err := kafka.InitConsumer(kconfig)
	if err != nil {
		panic(err)
	}
	wg.Add(1)
	go func() {
		for i := 0; i < 10000; i++ {
			err := produce.SendMessage(&sarama.ProducerMessage{
				Topic:     cfg.Kafka.WebTopic,
				Value:     sarama.StringEncoder(fmt.Sprintf("this is %d", i)),
				Partition: int32(i % 3),
			})
			if err != nil {
				fmt.Printf("第 %d 个消息生产失败\n", i)
			} else {
				fmt.Printf("第 %d 个消息生产成功\n", i)
			}
			time.Sleep(time.Second)
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for {
			consumer.RecvMessage(KafkacallBack, cfg.Kafka.WebTopic)
		}

	}()
	wg.Wait()
}
func KafkacallBack(message *sarama.ConsumerMessage) {
	data := map[string]string{
		"Partition": strconv.Itoa(int(message.Partition)),
		"Offset":    strconv.Itoa(int(message.Offset)),
		"Key":       string(message.Key),
		"Value":     string(message.Value),
	}
	fmt.Printf("%+v", data)
}
