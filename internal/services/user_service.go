package services

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// UserService is a service for managing users
type UserService struct {
	*baseService
}

// newUserService creates a new user service
func newUserService(bs *baseService) *UserService {
	return &UserService{
		baseService: bs,
	}
}

// GetAll returns all users
func (us *UserService) GetAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Write([]byte("getUsers GET /"))
	w.WriteHeader(http.StatusOK)
}

// GetByID returns a user by ID
func (us *UserService) GetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Write([]byte("getUserById GET /" + ps.ByName("id")))
	w.WriteHeader(http.StatusOK)
}

// DeleteByID deletes a user by ID
func (us *UserService) DeleteByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logoutUser := us.eventFactory.GetEvent("logout_user")
	s, _ := logoutUser.Call()
	w.Write([]byte(s + "\n"))
	w.Write([]byte("deleteUserById DELETE /" + ps.ByName("id") + "\n"))
	w.WriteHeader(http.StatusOK)
}
