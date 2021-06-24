package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nsqio/go-nsq"
)

func main() {
	topic := "test"
	channel := "test_channel"
	address := "127.0.0.1:4161"
	config := nsq.NewConfig()
	
	q, _ := nsq.NewConsumer(topic, channel, config)
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("接收到了訊息(字串)：%v", string(message.Body))

		return nil
	}))

	err := q.ConnectToNSQLookupd(address)
	if err != nil {
		log.Fatal(err)
	}
	
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)
	<- c

	// Gracefully stop
	q.Stop()
}
