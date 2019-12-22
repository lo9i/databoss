package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lo9i/databoss/server/domain"
	"log"
	"net/http"
)

type Server struct {
	Router           *mux.Router
	UserService      domain.UserService
	CandidateService domain.CandidateService
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	server.initializeRoutes()
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
