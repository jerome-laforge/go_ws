package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jerome-laforge/go_ws/dto"
	"os"
	"time"
)

var Repo repo

type repo struct {
	dbUrl string
	db    *sql.DB
}

func init() {
	host := "127.0.0.1"
	port := "3306"
	login := "todos_rw"
	passwd := "todos_rw"
	db := "todos"

	if tmp := os.Getenv("_MYSQL_HOST"); len(tmp) > 0 {
		host = tmp
	} else if tmp = os.Getenv("_ENV_MYSQL_HOST"); len(tmp) > 0 && len(os.Getenv(tmp)) > 0 {
		host = os.Getenv(tmp)
	}
	if tmp := os.Getenv("_MYSQL_PORT"); len(tmp) > 0 {
		port = tmp
	} else if tmp = os.Getenv("_ENV_MYSQL_PORT"); len(tmp) > 0 && len(os.Getenv(tmp)) > 0 {
		port = os.Getenv(tmp)
	}
	if tmp := os.Getenv("_MYSQL_LOGIN"); len(tmp) > 0 {
		login = tmp
	}
	if tmp := os.Getenv("_MYSQL_PASSWD"); len(tmp) > 0 {
		passwd = tmp
	}
	Repo.dbUrl = login + ":" + passwd + "@tcp(" + host + ":" + port + ")/" + db + "?parseTime=true"
	fmt.Println(Repo.dbUrl)

	var err error
	Repo.db, err = sql.Open("mysql", Repo.dbUrl)
	if err != nil {
		panic(err.Error())
	}
}

func (obj repo) RepoGetTodos() dto.Todos {
	rows, err := obj.db.Query("select id, name, completed, due from todos")
	if err != nil {
		panic(err.Error())
	}

	todos := make(dto.Todos, 0, 10)
	var curTodo dto.Todo
	for rows.Next() {
		err = rows.Scan(&curTodo.Id, &curTodo.Name, &curTodo.Completed, &curTodo.Due)
		if err != nil {
			panic(err)
		}
		todos = append(todos, curTodo)
	}
	return todos
}

func (obj repo) RepoCreateTodo(t dto.Todo) dto.Todo {
	if t.Due.IsZero() {
		t.Due = time.Now().Add(7 * 24 * time.Hour)
	}

	res, err := obj.db.Exec("insert into todos (name, completed, due) values (?, ?, ?)", t.Name, t.Completed, t.Due)
	if err != nil {
		panic(err)
	}

	id, _ := res.LastInsertId()

	t.Id = int(id)

	return t
}

func (obj repo) RepoDestroyTodo(id int) (dto.Todo, error) {
	if todo, ok := obj.RepoFindTodo(id); ok {
		_, err := obj.db.Exec("delete from todos where id = ?", id)
		if err != nil {
			return todo, err
		}
	}
	return dto.Todo{}, nil
}

func (obj repo) RepoFindTodo(id int) (dto.Todo, bool) {
	todo := new(dto.Todo)
	row := obj.db.QueryRow("select id, name, completed, due from todos where id = ?", id)
	err := row.Scan(&todo.Id, &todo.Name, &todo.Completed, &todo.Due)
	if err != nil {
		return *todo, false
	}

	return *todo, true
}
