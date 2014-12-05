package dao

import "database/sql"
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jerome-laforge/go_ws/dto"
	"os"
	"time"
	"fmt"
)

var host = "127.0.0.1"
var port = "3306"
var login = "todos_rw"
var passwd = "todos_rw"
var db = "todos"

var dbUrl string

func init() {
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
	dbUrl = login + ":" + passwd + "@tcp(" + host + ":" + port + ")/" + db + "?parseTime=true"
	fmt.Println(dbUrl)
}

func RepoGetTodos() dto.Todos {
	con, err := sql.Open("mysql", dbUrl)
	if err != nil {
		panic(err.Error())
	}
	defer con.Close()

	rows, err := con.Query("select id, name, completed, due from todos")
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

func RepoCreateTodo(t dto.Todo) dto.Todo {
	con, err := sql.Open("mysql", dbUrl)
	if err != nil {
		panic(err.Error())
	}
	defer con.Close()

	if t.Due.IsZero() {
		t.Due = time.Now().Add(7 * 24 * time.Hour)
	}

	res, err := con.Exec("insert into todos (name, completed, due) values (?, ?, ?)", t.Name, t.Completed, t.Due)
	if err != nil {
		panic(err)
	}

	id, _ := res.LastInsertId()

	t.Id = int(id)

	return t
}

func RepoDestroyTodo(id int) (dto.Todo, error) {
	con, err := sql.Open("mysql", dbUrl)
	if err != nil {
		panic(err.Error())
	}
	defer con.Close()

	if todo, ok := RepoFindTodo(id); ok {
		_, err = con.Exec("delete from todos where id = ?", id)
		if err != nil {
			return todo, err
		}
	}
	return dto.Todo{}, nil
}

func RepoFindTodo(id int) (dto.Todo, bool) {
	con, err := sql.Open("mysql", dbUrl)
	if err != nil {
		panic(err.Error())
	}
	defer con.Close()

	todo := new(dto.Todo)
	row := con.QueryRow("select id, name, completed, due from todos where id = ?", id)
	err = row.Scan(&todo.Id, &todo.Name, &todo.Completed, &todo.Due)
	if err != nil {
		return *todo, false
	}

	return *todo, true
}
