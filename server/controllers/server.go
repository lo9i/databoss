package controllers

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/lo9i/databoss/server/domain"
	"log"
	"net/http"
)

type Server struct {
	Router           *mux.Router
	UserService      domain.UserService
	CandidateService domain.CandidateService
	JobService       domain.JobService
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	server.initializeRoutes()
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	log.Fatal(http.ListenAndServe(addr, handlers.CORS(originsOk, headersOk, methodsOk)(server.Router)))
}
