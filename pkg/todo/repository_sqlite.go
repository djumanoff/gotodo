package todo

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	sqlite "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/l00p8/utils"
	_ "github.com/mattn/go-sqlite3"
	"strings"
)

type SqliteConfig struct {
	FilePath       string
	DbName         string
	MigrationsFile string
}

type sqliteDb struct {
	db *sql.DB
}

func NewSqliteRepository(cfg SqliteConfig) (Repository, error) {
	db, err := sql.Open("sqlite3", cfg.FilePath)
	if err != nil {
		return nil, err
	}
	driver, err := sqlite.WithInstance(db, &sqlite.Config{})
	if err != nil {
		return nil, err
	}
	m, err := migrate.NewWithDatabaseInstance("file://"+cfg.MigrationsFile, cfg.DbName, driver)
	if err != nil {
		return nil, err
	}
	_ = m.Steps(2)
	return &sqliteDb{db}, nil
}

func (db *sqliteDb) Update(upd *TodoUpdate) error {
	q := "UPDATE todos SET "
	var parts []string
	var values []interface{}
	if upd.Title != nil {
		parts = append(parts, "title = ?")
		values = append(values, *upd.Title)
	}
	if upd.Body != nil {
		parts = append(parts, "body = ?")
		values = append(values, *upd.Body)
	}
	if upd.Status != nil {
		parts = append(parts, "status = ?")
		values = append(values, *upd.Status)
	}
	if len(parts) == 0 {
		return ErrNothingToUpdate
	}
	q += strings.Join(parts, ", ")
	q += " WHERE id = ?"
	values = append(values, upd.ID)
	_, err := db.db.Exec(q, values...)
	if err != nil {
		return err
	}
	return nil
}

func (db *sqliteDb) Create(item *Todo) (int64, error) {
	ret, err := db.db.Exec("INSERT INTO todos (title, body, status, owner_id) VALUES (?, ?, ?, ?)", item.Title, item.Body, item.Status, item.OwnerID)
	if err != nil {
		return 0, err
	}
	id, err := ret.LastInsertId()
	if err != nil {
		return 0, err
	}
	if id == 0 {
		return 0, ErrTodoCreation
	}
	return id, nil
}

func (db *sqliteDb) FindAll(query *TodoQuery, params *utils.ListParams) ([]*Todo, error) {
	q := "SELECT id, title, body, status, owner_id FROM todos"
	var parts []string
	var values []interface{}
	if query.Status != "" {
		parts = append(parts, "status = ?")
		values = append(values, query.Status)
	}
	if len(parts) > 0 {
		q += " WHERE "
		q += strings.Join(parts, " AND ")
	}
	q += params.SQLOrderAndPaging()
	rows, err := db.db.Query(q, values...)
	if err != nil {
		return nil, err
	}
	var todos = make([]*Todo, 0, params.ItemsPerPage)
	for rows.Next() {
		item := &Todo{}
		err = rows.Scan(&item.ID, &item.Title, &item.Body, &item.Status, &item.OwnerID)
		if err != nil {
			return nil, err
		}
		todos = append(todos, item)
	}
	return todos, nil
}

// FindById(id int64) (*Todo, error)
func (db *sqliteDb) FindById(id int64) (*Todo, error) {
	item := &Todo{}
	item.ID = id
	err := db.db.QueryRow("SELECT title, body, status, owner_id FROM todos WHERE id = ?", id).
		Scan(&item.Title, &item.Body, &item.Status, &item.OwnerID)
	if err == sql.ErrNoRows {
		return nil, ErrTodoNotFound
	} else if err != nil {
		return nil, err
	}
	return item, nil
}

func (db *sqliteDb) Delete(id int64) error {
	ret, err := db.db.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		return err
	}
	err = rowsAffected(ret)
	if err != nil {
		return err
	}
	return nil
}

func (db *sqliteDb) Health() error {
	return db.db.Ping()
}

func rowsAffected(ret sql.Result) error {
	n, err := ret.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return ErrTodoNotFound
	}
	return nil
}
