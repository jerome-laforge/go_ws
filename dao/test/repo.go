package test

import (
	"fmt"
	"github.com/jerome-laforge/go_ws/dto"
	"sync"
	"time"
)

var FakeRepo repo

type repo struct {
	currentId int
	todos     dto.Todos
	lock      sync.RWMutex
}

// Give us some seed data
func init() {
	FakeRepo.RepoCreateTodo(dto.Todo{Name: "Write presentation", Due: time.Now()})
	FakeRepo.RepoCreateTodo(dto.Todo{Name: "Host meetup", Due: time.Date(2014, time.November, 13, 18, 30, 0, 0, time.UTC)})
}

func (t *repo) RepoGetTodos() dto.Todos {
	return t.todos
}

func (t *repo) RepoFindTodo(id int) (dto.Todo, bool) {
	t.lock.RLock()
	defer t.lock.RUnlock()

	for _, t := range t.todos {
		if t.Id == id {
			return t, true
		}
	}
	// return empty Todo if not found
	return dto.Todo{}, false
}

func (obj *repo) RepoCreateTodo(t dto.Todo) dto.Todo {
	obj.lock.Lock()
	defer obj.lock.Unlock()

	obj.currentId++
	t.Id = obj.currentId
	obj.todos = append(obj.todos, t)
	return t
}

func (obj *repo) RepoDestroyTodo(id int) (dto.Todo, error) {
	obj.lock.Lock()
	defer obj.lock.Unlock()

	for i, t := range obj.todos {
		if t.Id == id {
			obj.todos = append(obj.todos[:i], obj.todos[i+1:]...)
			return t, nil
		}
	}
	return dto.Todo{}, fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
