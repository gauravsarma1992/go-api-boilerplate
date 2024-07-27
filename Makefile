setup:
	mysql -uroot -e "CREATE DATABASE IF NOT EXISTS boilerplate_dev;" 

test:
	make setup
	go test -v .

build:
	make setup
	go build -o cmd/server.go

run:
	make setup
	go run cmd/server.go