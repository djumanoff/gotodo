package cqrses

type CommandHandler interface {
	Exec(cmd Command) ([]Event, interface{}, error)
}
