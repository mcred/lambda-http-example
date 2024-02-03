package services

import (
	"github.com/julienschmidt/httprouter"
	"github.com/mcred/lambda-http-example/internal/events"
	"net/http"
)

type BaseService struct {
	eventFactory *events.EventFactory
}

func NewBaseService(ef *events.EventFactory) *BaseService {
	return &BaseService{
		eventFactory: ef,
	}
}

type Service interface {
	GetAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	GetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	DeleteByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type ServiceFactory struct {
	services map[string]Service
}

func NewServiceFactory() *ServiceFactory {
	ef := events.GetEventFactory()
	bs := NewBaseService(ef)
	us := NewUserService(bs)

	return &ServiceFactory{
		services: map[string]Service{
			"UserService": us,
		},
	}
}

func (sf *ServiceFactory) GetService(name string) Service {
	return sf.services[name]
}
