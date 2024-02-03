package models

import (
	"github.com/julienschmidt/httprouter"
	"lightweight-route-framework/internal/events"
	"net/http"
)

type User struct {
	ID int `json:"id"`
}

func GetUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("getUsers GET /"))
	w.WriteHeader(http.StatusOK)
}

func GetUserById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Write([]byte("getUserById GET /" + ps.ByName("id")))
	w.WriteHeader(http.StatusOK)
}

func DeleteUserById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ef := events.GetEventFactory()
	logoutUser := ef.GetEvent("logout_user")
	s, _ := logoutUser.Call()
	w.Write([]byte(s + "\n"))
	w.Write([]byte("deleteUserById DELETE /" + ps.ByName("id") + "\n"))
	w.WriteHeader(http.StatusOK)
}
