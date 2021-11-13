package todo

import (
	"github.com/l00p8/utils"
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

func (repo *mockRepo) FindAll(query *TodoQuery, params *utils.ListParams) ([]*Todo, error) {
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
