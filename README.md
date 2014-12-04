go_ws
=====

This project is for testing purpose with WS/REST/JSON and golang.

To download this project and its dependencies :
Don't forget :
- to set your `GOPATH` env (for example : `mkdir -p ~/go/go_ws; export GOPATH=~/go/go_ws; export PATH=${GOPATH}/bin:${PATH}`)
- to set `http_proxy` (if you are behing proxy)
- `go get github.com/jerome-laforge/go_ws`

#If MySQL not install, download it with docker for example:
    docker pull tutum/mysql
    docker run -d -p 3306:3306 -e MYSQL_PASS="admin" tutum/mysql

#Mysql script
`mysql -uadmin -p"admin" -h127.0.0.1 -P3306` (or get password with `docker logs <contener_id>`)

    create database todos;

    CREATE USER 'todos_rw'@'%' IDENTIFIED BY 'todos_rw';
    flush privileges;
    GRANT ALL PRIVILEGES ON todos.* TO 'todos_rw' WITH GRANT OPTION;

    use todos;

    create table todos (
        id        mediumint not null auto_increment,
        name      char(255),
        completed boolean,
        due       timestamp,
        primary key (id)
    ) engine=myisam;

