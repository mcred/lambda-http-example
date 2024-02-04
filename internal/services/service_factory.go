package services

import (
	"github.com/julienschmidt/httprouter"
	"github.com/mcred/lambda-http-example/internal/events"
	"net/http"
)

// BaseService is a base service that all other services will inherit from
type BaseService struct {
	eventFactory *events.EventFactory
}

// NewBaseService creates a new base service
func NewBaseService(ef *events.EventFactory) *BaseService {
	return &BaseService{
		eventFactory: ef,
	}
}

// Service is an interface that all services will implement
type Service interface {
	GetAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	GetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	DeleteByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

// ServiceFactory is a factory for creating services
type ServiceFactory struct {
	services map[string]Service
}

// GetService returns a service by name
func (sf *ServiceFactory) GetService(name string) Service {
	return sf.services[name]
}

func (sf *ServiceFactory) RegisterService(name string, service Service) {
	sf.services[name] = service
}

// NewServiceFactory creates a new service factory
func NewServiceFactory() *ServiceFactory {
	ef := events.GetEventFactory()
	sf := &ServiceFactory{
		services: make(map[string]Service),
	}
	bs := NewBaseService(ef)
	sf.RegisterService("UserService", NewUserService(bs))
	return sf
}
