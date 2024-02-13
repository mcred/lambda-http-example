package services

import (
	"github.com/julienschmidt/httprouter"
	"github.com/mcred/lambda-http-example/internal/events"
	"net/http"
)

type baseService struct {
	eventFactory *events.EventFactory
}

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
	services map[ServiceName]Service
}

// GetService returns a service by name
func (sf *ServiceFactory) GetService(name ServiceName) Service {
	return sf.services[name]
}

func (sf *ServiceFactory) registerService(name ServiceName, service Service) {
	sf.services[name] = service
}

// NewServiceFactory creates a new service factory
func NewServiceFactory() *ServiceFactory {
	ef := events.NewEventFactory()
	sf := &ServiceFactory{
		services: make(map[ServiceName]Service),
	}
	bs := newBaseService(ef)
	sf.registerService(USER_SERVICE, newUserService(bs))
	return sf
}

// ServiceName type
type ServiceName int

const (
	USER_SERVICE ServiceName = iota
)
