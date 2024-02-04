package services

import (
	"github.com/julienschmidt/httprouter"
	"github.com/mcred/lambda-http-example/internal/events"
	"net/http"
)

// baseService is a base service that all other services will inherit from
type baseService struct {
	eventFactory *events.EventFactory
}

// newBaseService creates a new base service
func newBaseService(ef *events.EventFactory) *baseService {
	return &baseService{
		eventFactory: ef,
	}
}

// service is an interface that all services will implement
type service interface {
	GetAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	GetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	DeleteByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

// ServiceFactory is a factory for creating services
type ServiceFactory struct {
	services map[string]service
}

// GetService returns a service by name
func (sf *ServiceFactory) GetService(name string) service {
	return sf.services[name]
}

func (sf *ServiceFactory) registerService(name string, service service) {
	sf.services[name] = service
}

// NewServiceFactory creates a new service factory
func NewServiceFactory() *ServiceFactory {
	ef := events.GetEventFactory()
	sf := &ServiceFactory{
		services: make(map[string]service),
	}
	bs := newBaseService(ef)
	sf.registerService("UserService", newUserService(bs))
	return sf
}
