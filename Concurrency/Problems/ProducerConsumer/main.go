package main

import (
	"fmt"
	"sync"
	"time"
)

const MAX_MSG_CHANNEL_SIZE = 10

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	cond := sync.NewCond(&mu)

	msgChannel := NewMsgChannel(MAX_MSG_CHANNEL_SIZE)
	producer := NewProducer(cond, msgChannel)
	consumer := NewConsumer(cond, msgChannel)
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := range 20 {
			producer.Produce(fmt.Sprintf("Message %d", i))
		}
	}()

	go func() {
		defer wg.Done()
		for range 20 {
			consumer.Consume()
		}
	}()
	wg.Wait()
}

type MsgChannel struct {
	maxBufferSize int
	buffer        []string
}

func NewMsgChannel(size int) *MsgChannel {
	return &MsgChannel{
		maxBufferSize: size,
		buffer:        make([]string, 0, size),
	}
}

func (mc *MsgChannel) IsEmpty() bool {
	return len(mc.buffer) == 0
}

func (mc *MsgChannel) IsFull() bool {
	return len(mc.buffer) == mc.maxBufferSize
}

func (mc *MsgChannel) Add(msg string) {
	mc.buffer = append(mc.buffer, msg)
}

func (mc *MsgChannel) Get() string {
	msg := mc.buffer[0]
	mc.buffer = mc.buffer[1:]
	return msg
}

type Producer struct {
	cond       *sync.Cond
	msgChannel *MsgChannel
}

func NewProducer(cond *sync.Cond, msgChannel *MsgChannel) *Producer {
	return &Producer{
		cond:       cond,
		msgChannel: msgChannel,
	}
}

func (p *Producer) Produce(msg string) {
	time.Sleep(500 * time.Millisecond)

	p.cond.L.Lock()
	defer p.cond.L.Unlock()

	for p.msgChannel.IsFull() {
		fmt.Println("Producer is waiting because msg channel is full")
		p.cond.Wait()
	}

	p.msgChannel.Add(msg)
	fmt.Println("Producer produced the msg:", msg)

	p.cond.Signal()
}

type Consumer struct {
	id         int
	cond       *sync.Cond
	msgChannel *MsgChannel
}

func NewConsumer(cond *sync.Cond, msgChannel *MsgChannel) *Consumer {
	return &Consumer{
		cond:       cond,
		msgChannel: msgChannel,
	}
}

func (c *Consumer) Consume() {
	time.Sleep(1 * time.Second)

	c.cond.L.Lock()
	defer c.cond.L.Unlock()

	for c.msgChannel.IsEmpty() {
		fmt.Println("Consumer is waiting because msg channel is empty")
		c.cond.Wait()
	}

	msg := c.msgChannel.Get()
	fmt.Println("Consumer consumed the msg: ", msg)
	c.cond.Signal()

}
