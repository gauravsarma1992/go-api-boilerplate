package main

import (
	"log"

	"github.com/go-api-boilerplate/server"
)

func main() {
	var (
		srv *server.ApiServer
		err error
	)
	if srv, err = server.NewServer(nil); err != nil {
		log.Fatal(err)
	}
	if err = srv.Run(); err != nil {
		log.Fatal(err)
	}
}
