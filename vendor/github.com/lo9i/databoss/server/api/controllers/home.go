package controllers

import (
	responses2 "github.com/lo9i/databoss/server/responses"
	"net/http"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses2.JSON(w, http.StatusOK, "Welcome To This Awesome API")
}