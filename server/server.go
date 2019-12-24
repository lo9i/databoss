package main

import (
	"github.com/gorilla/mux"
	"github.com/lo9i/databoss/server/controllers"
	"github.com/lo9i/databoss/server/domain"
)

func NewServer(uService domain.UserService,
				cService domain.CandidateService,
				jService domain.JobService) *controllers.Server {
	return &controllers.Server{
		UserService:      uService,
		CandidateService: cService,
		JobService: jService,
		Router:           mux.NewRouter(),
	}
}
