go_ws
=====

This project is for testing purpose with WS/REST/JSON and golang.

#To download this project and its dependencies :
Do steps below :
- to set your `GOPATH` env (for example : `mkdir -p ~/go/go_ws; export GOPATH=~/go/go_ws; export PATH=${GOPATH}/bin:${PATH}`)
- to set `http_proxy` (if you are behing proxy)
- `go get github.com/jerome-laforge/go_ws`

#If MySQL not install, download it with docker for example:
    docker pull mysql
    docker run --name some-mysql -e MYSQL_ROOT_PASSWORD=mysecretpassword -d mysql

#Mysql script
`docker run -it --link some-mysql:mysql --rm mysql sh -c 'exec mysql -h"$MYSQL_PORT_3306_TCP_ADDR" -P"$MYSQL_PORT_3306_TCP_PORT" -uroot -p"$MYSQL_ENV_MYSQL_ROOT_PASSWORD"'`

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

#Create docker's image
    go install github.com/jerome-laforge/go_ws
    echo 'FROM scratch:latest
    ADD go_ws /go_ws
    CMD [ "/go_ws" ]' > ${GOPATH}/bin/Dockerfile
    docker build -t go_ws ${GOPATH}/bin/

    docker run -d --link some-mysql:mysql -p 8080:8080 -e _ENV_MYSQL_HOST="MYSQL_PORT_3306_TCP_ADDR"  go_ws

#Unit Test
    go test -cover
    go test -coverprofile=/tmp/cover.out
    go tool cover -html=/tmp/cover.out
