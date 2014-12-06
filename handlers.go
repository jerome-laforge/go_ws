package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jerome-laforge/go_ws/dao"
	"github.com/jerome-laforge/go_ws/dao/mysql"
	"github.com/jerome-laforge/go_ws/dto"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

var _dao dao.Repo = mysql.Repo

func SetDao(dao dao.Repo) {
	_dao = dao
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(_dao.RepoGetTodos()); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var todoId int
	var err error
	if todoId, err = strconv.Atoi(vars["todoId"]); err != nil {
		panic(err)
	}

	if todo, ok := _dao.RepoFindTodo(todoId); ok {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(todo); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}

/*
Test with this curl command:

curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/todos

*/
func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo dto.Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := _dao.RepoCreateTodo(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

/*
Test with this curl command:

curl  -H "Content-Type: application/json"  -X DELETE http://localhost:8080/todos/2

*/

func TodoDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars != nil {
		if sTodoId, ok := vars["todoId"]; ok {
			var todo dto.Todo
			todoId, err := strconv.Atoi(sTodoId)
			if err == nil {
				todo, err = _dao.RepoDestroyTodo(todoId)
			}
			if err != nil {
				panic(err)
			} else {
				if err = json.NewEncoder(w).Encode(todo); err != nil {
					panic(err)
				}
			}
		}
	}
}
