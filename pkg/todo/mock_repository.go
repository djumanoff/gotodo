package todo

import http_helper "github.com/djumanoff/gotodo/pkg/http-helper"

func NewMockRepo() Repository {
	return &mockRepo{}
}

type mockRepo struct{}

func (repo *mockRepo) Update(upd *TodoUpdate) error {
	return nil
}

func (repo *mockRepo) Create(todo *Todo) (int64, error) {
	return 0, nil
}

func (repo *mockRepo) FindAll(params *http_helper.ListParams) ([]*Todo, error) {
	todos := []*Todo{}
	return todos, nil
}

func (repo *mockRepo) FindById(id int64) (*Todo, error) {
	return &Todo{}, nil
}

func (repo *mockRepo) Delete(id int64) error {
	return nil
}
