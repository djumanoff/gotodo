package cqrses

type Command interface {
	Exec() ([]Event, interface{}, error)
}
