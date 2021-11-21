package todo

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/l00p8/cqrses"
	"github.com/l00p8/l00p8"
	"io/ioutil"
	"net/http"
	"strconv"
)

type HttpHandlerFactory interface {
	CreateTodo() l00p8.Handler

	GetTodos() l00p8.Handler

	GetTodo(idParam string) l00p8.Handler
}

func NewHttpHandlerFactory(cmder cqrses.CommandHandler, errSys l00p8.ErrorSystem) HttpHandlerFactory {
	return &httpHandlerFactory{cmder, errSys}
}

type httpHandlerFactory struct {
	cmder  cqrses.CommandHandler
	errSys l00p8.ErrorSystem
}

func (fac *httpHandlerFactory) CreateTodo() l00p8.Handler {
	return func(r *http.Request) l00p8.Response {
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
		return l00p8.Created(todo)
	}
}

func (fac *httpHandlerFactory) GetTodos() l00p8.Handler {
	return func(r *http.Request) l00p8.Response {
		cmd := &CommandGetTodos{}
		query := &TodoQuery{}
		lparams := l00p8.ParseFromRequest(r, query)
		cmd.Query = query
		cmd.ListParams = lparams
		_, todos, err := fac.cmder.Exec(cmd)
		if err != nil {
			return fac.errSys.BadRequest(30, err.Error())
		}
		return l00p8.OK(todos)
	}
}

func (fac *httpHandlerFactory) GetTodo(idParam string) l00p8.Handler {
	return func(r *http.Request) l00p8.Response {
		idStr := chi.URLParam(r, idParam)
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return fac.errSys.BadRequest(20, err.Error())
		}
		cmd := &CommandGetTodoByID{ID: id}
		_, todos, err := fac.cmder.Exec(cmd)
		if err == ErrTodoNotFound {
			return fac.errSys.NotFound(30, err.Error())
		} else if err != nil {
			return fac.errSys.BadRequest(40, err.Error())
		}
		return l00p8.OK(todos)
	}
}
