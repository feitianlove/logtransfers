package main

import (
	"fmt"
	"github.com/feitianlove/logtransfers/config"
)

func main() {
	cfg, err := config.NewConfig("./etc/logtransfer.conf")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", cfg.Kafka)
}
