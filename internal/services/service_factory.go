package services

import (
	"github.com/julienschmidt/httprouter"
	"github.com/mcred/lambda-http-example/internal/events"
	"net/http"
)

// baseService is a base Service that all other services will inherit from
type baseService struct {
	eventFactory *events.EventFactory
}

// newBaseService creates a new base Service
func newBaseService(ef *events.EventFactory) *baseService {
	return &baseService{
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

func (sf *ServiceFactory) registerService(name string, service Service) {
	sf.services[name] = service
}

// NewServiceFactory creates a new service factory
func NewServiceFactory() *ServiceFactory {
	ef := events.NewEventFactory()
	sf := &ServiceFactory{
		services: make(map[string]Service),
	}
	bs := newBaseService(ef)
	sf.registerService("UserService", newUserService(bs))
	return sf
}
