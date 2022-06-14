package queue

type (
	Publisher interface {
		publish(nameQueue string, body string) error
	}

	Consumer interface {
		consume(fn func(string) string, nameQueue string)
	}

	Queue interface {
		Publisher
		Consumer
	}

	Broker struct {
		Queue
	}
)

func NewQueueBroker(q Queue) *Broker {
	return &Broker{
		Queue: q,
	}
}

func (b Broker) PublishMessage(nameQueue string, message string) error {
	err := b.Queue.publish(nameQueue, message)
	if err != nil {
		return err
	}
	return nil
}

func (b Broker) ConsumeMessage(fn func(string) string, nameQueue string) {
	b.consume(fn, nameQueue)
}
