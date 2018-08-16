package main

import (
	"sync"
	"github.com/Shopify/sarama"
	"strings"
	"fmt"
)

var (
	wg     sync.WaitGroup
)

func main() {
	consumer, err := sarama.NewConsumer(strings.Split("localhost:9092", ","), nil)
	if err != nil {
		fmt.Println("Failed to start consumer:", err)
	}
	partitionList, err := consumer.Partitions("test")
	if err != nil {
		fmt.Println("Failed to get the list of partitions: ", err)
	}
	fmt.Println("par:",partitionList)
	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("test", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("Failed to start consumer for partition %d: %s\n", partition, err)
		}
		wg.Add(1)
		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			defer pc.AsyncClose()
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				fmt.Println()
			}
		}(pc)
	}
	wg.Wait()
	fmt.Println("Done consuming topic hello")
	consumer.Close()
}