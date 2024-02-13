package services

import (
	"github.com/julienschmidt/httprouter"
	"github.com/mcred/lambda-http-example/internal/events"
	"net/http"
)

// UserService is a service for managing users
type UserService struct {
	*baseService
}

// newUserService creates a new user Service
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
	s, _ := us.eventFactory.Call(events.LOGOUT_USER)
	w.Write([]byte(s + "\n"))
	w.Write([]byte("deleteUserById DELETE /" + ps.ByName("id") + "\n"))
	w.WriteHeader(http.StatusOK)
}
