package cqrses

type EventPublisher interface {
	Publish(Event) error
}

type (
	kafkaPublisher struct{}

	kafkaConfig struct{}
)

func NewKafkaPublisher(cfg kafkaConfig) EventPublisher {
	return &kafkaPublisher{}
}

func (pub *kafkaPublisher) Publish(Event) error {
	return nil
}
