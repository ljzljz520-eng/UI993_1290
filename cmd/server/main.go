package main

import (
	"log"
	"net/http"
	"refereequal/admin"
	"refereequal/query"
	"refereequal/repository"
	"refereequal/storage"
	"refereequal/web"
)

func main() {
	s, e := storage.Open("referees.db")
	if e != nil {
		log.Fatal(e)
	}
	defer s.Close()
	r := repository.New(s)
	_ = admin.New(r)
	q := query.New(r)
	log.Fatal(http.ListenAndServe(":8080", web.New(q).Handler()))
}
