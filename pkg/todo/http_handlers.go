package todo

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/l00p8/cqrses"
	server "github.com/l00p8/http-server"
	"github.com/l00p8/utils"
	"io/ioutil"
	"net/http"
	"strconv"
)

type HttpHandlerFactory interface {
	CreateTodo() server.Handler

	GetTodos() server.Handler

	GetTodo(idParam string) server.Handler
}

func NewHttpHandlerFactory(cmder cqrses.CommandHandler, errSys utils.ErrorSystem) HttpHandlerFactory {
	return &httpHandlerFactory{cmder, errSys}
}

type httpHandlerFactory struct {
	cmder  cqrses.CommandHandler
	errSys utils.ErrorSystem
}

func (fac *httpHandlerFactory) CreateTodo() server.Handler {
	return func(r *http.Request) utils.Response {
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
		return utils.Created(todo)
	}
}

func (fac *httpHandlerFactory) GetTodos() server.Handler {
	return func(r *http.Request) utils.Response {
		cmd := &CommandGetTodos{}
		query := &TodoQuery{}
		lparams := utils.ParseFromRequest(r, query)
		cmd.Query = query
		cmd.ListParams = lparams
		_, todos, err := fac.cmder.Exec(cmd)
		if err != nil {
			return fac.errSys.BadRequest(30, err.Error())
		}
		return utils.OK(todos)
	}
}

func (fac *httpHandlerFactory) GetTodo(idParam string) server.Handler {
	return func(r *http.Request) utils.Response {
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
		return utils.OK(todos)
	}
}
