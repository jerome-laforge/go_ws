go_ws
=====

This project is for testing purpose with WS/REST/JSON and golang.

#To download this project and its dependencies :
Do steps below :
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

#Cross compile
amd64 (a.k.a. x86-64); 6g,6l,6c,6a : A mature implementation. The compiler has an effective optimizer (registerizer) and generates good code (although gccgo can do noticeably better sometimes). 
- `GOOS=windows GOARCH=amd64 go install github.com/jerome-laforge/go_ws`

386 (a.k.a. x86 or x86-32); 8g,8l,8c,8a : Comparable to the amd64 port. 
- `GOOS=windows GOARCH=386 go install github.com/jerome-laforge/go_ws`
 
arm (a.k.a. ARM); 5g,5l,5c,5a : Supports Linux, FreeBSD and NetBSD binaries. Less widely used than the other ports. 
- `GOOS=windows GOARCH=arm go install github.com/jerome-laforge/go_ws`

 
$GOOS 	$GOARCH
-	darwin 	386
-	darwin 	amd64
-	dragonfly 	386
-	dragonfly 	amd64
-	freebsd 	386
-	freebsd 	amd64
-	freebsd 	arm
-	linux 	386
-	linux 	amd64
-	linux 	arm
-	netbsd 	386
-	netbsd 	amd64
-	netbsd 	arm
-	openbsd 	386
-	openbsd 	amd64
-	plan9 	386
-	plan9 	amd64
-	solaris 	amd64
-	windows 	386
-	windows 	amd64
