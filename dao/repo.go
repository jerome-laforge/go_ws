package dao

import "github.com/jerome-laforge/go_ws/dto"

type Repo interface {
	RepoGetTodos() dto.Todos
	RepoCreateTodo(t dto.Todo) dto.Todo
	RepoDestroyTodo(id int) (dto.Todo, error)
	RepoFindTodo(id int) (dto.Todo, bool)
}
