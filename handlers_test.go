package go_ws

import (
	"github.com/jerome-laforge/go_ws/dao/test"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func init() {
	Dao = &test.FakeRepo
}

func TestTodoCreate(t *testing.T) {
	test.FakeRepo.Clear()
	if len(test.FakeRepo.RepoGetTodos()) != 0 {
		t.Fatalf("repository is not empty")
	}
	r, err := http.NewRequest("POST", "/todos", strings.NewReader(`{"name":"New Todo 0"}`))
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	TodoCreate(w, r)
	if w.Code != http.StatusCreated {
		t.Fatalf("Bad status code [expected: %d] [actual: %d]", http.StatusCreated, w.Code)
	}
	if len(test.FakeRepo.RepoGetTodos()) != 1 {
		t.Fatalf("repository length incositency")
	}
}

func TestTodoCreateBad(t *testing.T) {
	test.FakeRepo.Clear()
	if len(test.FakeRepo.RepoGetTodos()) != 0 {
		t.Fatalf("repository is not empty")
	}
	r, err := http.NewRequest("POST", "/todos", strings.NewReader(`{"name":"New Todo 0"`))
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	TodoCreate(w, r)
	if w.Code != 422 {
		t.Fatalf("Bad status code [expected: %d] [actual: %d]", 422, w.Code)
	}
	if len(test.FakeRepo.RepoGetTodos()) != 1 {
		t.Fatalf("repository length incositency")
	}
}

func TestIndex(t *testing.T) {
	test.FakeRepo.Clear()
	w := httptest.NewRecorder()
	Index(w, nil)
	if w.Body.String() != "Welcome!\n" {
		t.Fail()
	}
}

func TestTodoIndex(_ *testing.T) {
	test.FakeRepo.Clear()
	test.FakeRepo.Init()
	w := httptest.NewRecorder()
	TodoIndex(w, nil)
}
