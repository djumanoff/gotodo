package todo

import (
	"context"
	"encoding/json"
	"github.com/djumanoff/gotodo/pkg/cqrses"
	http_helper "github.com/djumanoff/gotodo/pkg/http-helper"
	"io/ioutil"
	"net/http"
)

type HttpEndpointFactory struct {
	cmder  cqrses.CommandHandler
	errSys http_helper.ErrorSystem
}

func (fac *HttpEndpointFactory) CreateTodo() http_helper.Handler {
	return func(ctx context.Context, r *http.Request) http_helper.Response {
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
