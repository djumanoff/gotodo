package todo

import (
	"encoding/json"
	"github.com/djumanoff/gotodo/pkg/cqrses"
	http_helper "github.com/djumanoff/gotodo/pkg/http-helper"
	"io/ioutil"
	"net/http"
)

type HttpHandlerFactory interface {
	CreateTodo() http_helper.Handler

	GetTodos() http_helper.Handler
}

func NewHttpHandlerFactory(cmder cqrses.CommandHandler, errSys http_helper.ErrorSystem) HttpHandlerFactory {
	return &httpHandlerFactory{cmder, errSys}
}

type httpHandlerFactory struct {
	cmder  cqrses.CommandHandler
	errSys http_helper.ErrorSystem
}

func (fac *httpHandlerFactory) CreateTodo() http_helper.Handler {
	return func(r *http.Request) http_helper.Response {
		cmd := &CommandCreateTodo{}
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return fac.errSys.BadRequest(20, err.Error())
		}
		err = json.Unmarshal(d, cmd)
		if err != nil {
			return fac.errSys.BadRequest(30, err.Error())
		}
		_, todo, err := fac.cmder.Exec(cmd)
		if err != nil {
			return fac.errSys.BadRequest(30, err.Error())
		}
		return http_helper.Created(todo)
	}
}

func (fac *httpHandlerFactory) GetTodos() http_helper.Handler {
	return func(r *http.Request) http_helper.Response {
		cmd := &CommandGetTodos{}
		_, todos, err := fac.cmder.Exec(cmd)
		if err != nil {
			return fac.errSys.BadRequest(30, err.Error())
		}
		return http_helper.OK(todos)
	}
}
