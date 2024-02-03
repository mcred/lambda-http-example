package services

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type UserService struct {
	*BaseService
}

func NewUserService(bs *BaseService) *UserService {
	return &UserService{
		BaseService: bs,
	}
}

func (us *UserService) GetAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Write([]byte("getUsers GET /"))
	w.WriteHeader(http.StatusOK)
}

func (us *UserService) GetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Write([]byte("getUserById GET /" + ps.ByName("id")))
	w.WriteHeader(http.StatusOK)
}

func (us *UserService) DeleteByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logoutUser := us.eventFactory.GetEvent("logout_user")
	s, _ := logoutUser.Call()
	w.Write([]byte(s + "\n"))
	w.Write([]byte("deleteUserById DELETE /" + ps.ByName("id") + "\n"))
	w.WriteHeader(http.StatusOK)
}
