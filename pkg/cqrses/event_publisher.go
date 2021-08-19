package cqrses

type EventPublisher interface {
	Publish(Event) error
}

type (
	kafkaPublisher struct{}

	KafkaConfig struct{}
)

func NewKafkaPublisher(cfg KafkaConfig) EventPublisher {
	return &kafkaPublisher{}
}

func (pub *kafkaPublisher) Publish(Event) error {
	return nil
}
