package cqrses

type CommandHandler interface {
	Exec(cmd Command) ([]Event, interface{}, error)
}

func NewCommandHandler(svc interface{}) CommandHandler {
	return &cmder{svc}
}

type cmder struct {
	svc interface{}
}

func (c *cmder) Exec(cmd Command) ([]Event, interface{}, error) {
	return cmd.Exec(c.svc)
}
