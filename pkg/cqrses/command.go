package cqrses

type Command interface {
	Exec(svc interface{}) ([]Event, interface{}, error)
}
