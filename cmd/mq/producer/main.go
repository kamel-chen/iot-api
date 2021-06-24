package main

import (
	"log"

	"github.com/nsqio/go-nsq"
)

func main() {
	config := nsq.NewConfig()

	p, _ := nsq.NewProducer("127.0.0.1:4150", config)

	err := p.Publish("test", []byte("hello world!"))
	if err != nil {
		log.Fatal(err)
	}

	// Gracefully stop
	p.Stop()
}
