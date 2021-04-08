package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	goliblogger "github.com/feitianlove/golib/common/logger"
	ftK "github.com/feitianlove/golib/kafka"
	"github.com/feitianlove/logtransfers/config"
	"github.com/feitianlove/logtransfers/kafka"
	"github.com/feitianlove/logtransfers/logger"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

//TODO golib 中初始化kafka变成 string
func main() {
	var wg sync.WaitGroup
	cfg, err := config.NewConfig("../etc/logtransfer.conf")
	if err != nil {
		panic(err)
	}
	fmt.Printf("config init success: %+v", cfg.Kafka)

	//初始化logger
	err = logger.InitCtrlLog(&goliblogger.LogConf{
		LogLevel:      "info",
		LogPath:       "../log/logtransfer",
		LogReserveDay: 7,
		ReportCaller:  false,
	})
	if err != nil {
		panic(err)
	}
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
				logger.Ctrl.WithFields(logrus.Fields{
					"Partition": int32(i % 3),
					"topic":     cfg.Kafka.WebTopic,
					"Value":     fmt.Sprintf("第 %d 个消息生产失败", i),
				}).Info("product message")
			} else {
				logger.Ctrl.WithFields(logrus.Fields{
					"Partition": int32(i % 3),
					"topic":     cfg.Kafka.WebTopic,
					"Value":     fmt.Sprintf("第 %d 个消息生产成功", i),
				}).Info("SendMessage")
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
	logger.Ctrl.WithFields(logrus.Fields{
		"Partition": message.Partition,
		"Offset":    message.Offset,
		"Key":       string(message.Key),
		"Value":     string(message.Value),
	}).Info("RecvMessage")
}
