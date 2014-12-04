package dao

import (
	"fmt"
	"sync"
	"time"
	"github.com/jerome-laforge/go_ws/dto"
)

var currentId int = 0

var todos dto.Todos

var lock sync.RWMutex

// Give us some seed data
func init() {
	RepoCreateTodo(dto.Todo{Name: "Write presentation", Due: time.Now()})
	RepoCreateTodo(dto.Todo{Name: "Host meetup", Due: time.Date(2014, time.November, 13, 18, 30, 0, 0, time.UTC)})
}

func RepoGetTodos() dto.Todos {
	return todos
}

func RepoFindTodo(id int) (dto.Todo, bool) {
	lock.RLock()
	defer lock.RUnlock()

	for _, t := range todos {
		if t.Id == id {
			return t, true
		}
	}
	// return empty Todo if not found
	return dto.Todo{}, false
}

func RepoCreateTodo(t dto.Todo) dto.Todo {
	lock.Lock()
	defer lock.Unlock()

	currentId++
	t.Id = currentId
	todos = append(todos, t)
	return t
}

func RepoDestroyTodo(id int) (dto.Todo, error) {
	lock.Lock()
	defer lock.Unlock()

	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return t, nil
		}
	}
	return dto.Todo{}, fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
