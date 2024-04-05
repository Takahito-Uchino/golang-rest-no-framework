package model

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type TodoModel struct {
	db *sql.DB
}

func NewTodoModel() *TodoModel {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=todo_db sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	return &TodoModel{db}
}

type Todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Completed bool   `json:"completed"`
}

func (t *TodoModel) GetTodos() (todos []Todo, err error) {
	query := "SELECT id, title, content, completed FROM todos"
	rows, err := t.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos = []Todo{}
	for rows.Next() {
		todo := Todo{}
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Content, &todo.Completed)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return todos, nil
}

func (t *TodoModel) CreateTodo(todo *Todo) (err error) {
	query := "INSERT INTO todos (title, content, completed) VALUES ($1, $2, $3)"
	result, err := t.db.Exec(query, todo.Title, todo.Content, todo.Completed)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	todo.Id = int(id)

	return nil
}

func (t *TodoModel) UpdateTodo(id int, todo *Todo) (err error) {
	query := "UPDATE todos SET title = $1, content = $2, completed = $3 WHERE id = $4"
	_, err = t.db.Exec(query, todo.Title, todo.Content, todo.Completed, id)
	if err != nil {
		return err
	}
	return nil
}

func (t *TodoModel) DeleteTodo(id int) (err error) {
	query := "DELETE FROM todos WHERE id = $1"
	_, err = t.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
