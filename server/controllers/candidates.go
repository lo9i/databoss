package controllers

import (
	"encoding/json"
	"errors"
	"github.com/lo9i/databoss/server/auth"
	"github.com/lo9i/databoss/server/domain"
	"github.com/lo9i/databoss/server/responses"
	"github.com/lo9i/databoss/server/utils/formaterror"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//
//func (server *Server) CreateCandidate(w http.ResponseWriter, r *http.Request) {
//
//	body, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		responses.ERROR(w, http.StatusUnprocessableEntity, err)
//	}
//	candidate := domain.Candidate{}
//	err = json.Unmarshal(body, &candidate)
//	if err != nil {
//		responses.ERROR(w, http.StatusUnprocessableEntity, err)
//		return
//	}
//	//candidate.Prepare(
//	//err = candidate.Validate("")
//	//if err != nil {
//	//	responses.ERROR(w, http.StatusUnprocessableEntity, err)
//	//	return
//	//}
//	candidateCreated, err := server.CandidateService.Save(&candidate)
//
//	if err != nil {
//
//		formattedError := formaterror.FormatError(err.Error())
//
//		responses.ERROR(w, http.StatusInternalServerError, formattedError)
//		return
//	}
//	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, candidateCreated.Id))
//	responses.JSON(w, http.StatusCreated, candidateCreated)
//}

func (server *Server) GetCandidates(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	var candidates *[]domain.Candidate
	_, err := strconv.ParseUint(userId, 10, 32)
	if err != nil {
		candidates = server.CandidateService.Find()
	} else {
		candidates = server.CandidateService.Find("userId = ?", userId)
	}
	responses.JSON(w, http.StatusOK, candidates)
}

func (server *Server) GetCandidateById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	candidate := server.CandidateService.Get(uid)
	responses.JSON(w, http.StatusOK, candidate)
}

func (server *Server) GetCandidateByUserId(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	candidate := server.CandidateService.GetByUserId(uid)
	responses.JSON(w, http.StatusOK, candidate)
}

func (server *Server) UpdateCandidate(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	candidate := domain.Candidate{}
	err = json.Unmarshal(body, &candidate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if tokenID != uint32(uid) {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	//candidate.Prepare()
	//err = candidate.Validate("update")
	//if err != nil {
	//	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	//	return
	//}
	updatedCandidate, err := server.CandidateService.Save(&candidate)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedCandidate)
}

//func (server *Server) DeleteCandidate(w http.ResponseWriter, r *http.Request) {
//
//	vars := mux.Vars(r)
//
//	candidate := domain.Candidate{}
//
//	uid, err := strconv.ParseUint(vars["id"], 10, 32)
//	if err != nil {
//		responses.ERROR(w, http.StatusBadRequest, err)
//		return
//	}
//	tokenID, err := auth.ExtractTokenID(r)
//	if err != nil {
//		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
//		return
//	}
//	if tokenID != 0 && tokenID != uint32(uid) {
//		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
//		return
//	}
//	_, err = candidate.DeleteACandidate(server.DB, uint32(uid))
//	if err != nil {
//		responses.ERROR(w, http.StatusInternalServerError, err)
//		return
//	}
//	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
//	responses.JSON(w, http.StatusNoContent, "")
//}
