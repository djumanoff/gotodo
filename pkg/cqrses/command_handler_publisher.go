package cqrses

func NewPublishingCommandHandler(pub EventPublisher, cmder CommandHandler) CommandHandler {
	return &pubCmder{pub, cmder}
}

type pubCmder struct {
	pub   EventPublisher
	cmder CommandHandler
}

func (c *pubCmder) Exec(cmd Command) ([]Event, interface{}, error) {
	events, data, err := c.cmder.Exec(cmd)
	for _, event := range events {
		// ignore errors for now
		_ = c.pub.Publish(event)
	}
	return events, data, err
}
