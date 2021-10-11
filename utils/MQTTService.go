package utils

import (
	"github.com/golang/glog"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func GetProducer() (*kafka.Producer){
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "10.16.150.138:6667"})
	if err != nil {
		panic(err)
	}
	return p
}
func PublishMessage( value []byte,topic string,p *kafka.Producer){
	err:=p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:         value,
		Key: []byte("key_"+topic),
	}, nil)
	if err!=nil{
		glog.Error(err)
	}
}

