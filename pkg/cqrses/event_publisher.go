package cqrses

type EventPublisher interface {
	Publish(Event) error
}

type kafkaPublisher struct {
}
