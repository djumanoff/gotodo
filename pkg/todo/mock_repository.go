package todo

import (
	http_helper "github.com/djumanoff/gotodo/pkg/utils"
)

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
	var todos []*Todo
	return todos, nil
}

func (repo *mockRepo) FindById(id int64) (*Todo, error) {
	return &Todo{}, nil
}

func (repo *mockRepo) Delete(id int64) error {
	return nil
}

func (repo *mockRepo) Health() error {
	return nil
}
