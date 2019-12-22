package controllers

import (
	//"encoding/json"
	//"github.com/lo9i/databoss/server/api/auth"
	//"github.com/lo9i/databoss/server/api/domain"
	//"github.com/lo9i/databoss/server/api/responses"
	//"github.com/lo9i/databoss/server/api/utils/formaterror"
	//"io/ioutil"
	"net/http"

	//"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	//body, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	//	return
	//}
	//user := domain.User{}
	//err = json.Unmarshal(body, &user)
	//if err != nil {
	//	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	//	return
	//}
	//
	//user.Prepare()
	//err = user.Validate("login")
	//if err != nil {
	//	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	//	return
	//}
	//token, err := server.SignIn(user.Email, user.Password)
	//if err != nil {
	//	formattedError := formaterror.FormatError(err.Error())
	//	responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
	//	return
	//}
	//responses.JSON(w, http.StatusOK, token)
}

func (server *Server) SignIn(email, password string) (string, error) {

	//var err error

	//user := domain.User{}
	//
	//err = server.DB.Debug().Model(domain.User{}).Where("email = ?", email).Take(&user).Error
	//if err != nil {
	//	return "", err
	//}
	//err = domain.VerifyPassword(user.Password, password)
	//if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
	//	return "", err
	//}
	//return auth.CreateToken(user.ID)
}
