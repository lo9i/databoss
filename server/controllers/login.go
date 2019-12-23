package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/lo9i/databoss/server/auth"
	"github.com/lo9i/databoss/server/domain"
	"github.com/lo9i/databoss/server/responses"
	"github.com/lo9i/databoss/server/utils/formaterror"
	"io/ioutil"
	"net/http"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := domain.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(user.Username, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) SignIn(email, password string) (map[string]string, error) {

	//var err error
	user := server.UserService.FindUserByEmail(email)
	if user == nil {
		return nil, fmt.Errorf("Verifique las credenciales provistas.")
	}

	//err = domain.VerifyPassword(user.Password, password)
	//if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
	//	return "", err
	//}
	return auth.CreateToken(user.Id)
}

func (server *Server) Logout(w http.ResponseWriter, r *http.Request) {
	//body, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	//	return
	//}

	//token := body
	//server.SignOut(token)
	//responses.JSON(w, http.StatusOK)
	responses.JSON(w, http.StatusOK, "")
}

func (server *Server) Refresh(w http.ResponseWriter, r *http.Request) {
	token, err := auth.RefreshToken(r)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}
