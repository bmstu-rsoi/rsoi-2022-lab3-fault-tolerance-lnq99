package services

import (
	"github.com/gofiber/fiber/v2"
)

const (
	DefaultIP = "localhost"
)

var (
	ApiVersion      = "api/v1"
	BonusServiceIP  = ""
	FlightServiceIP = ""
	TicketServiceIP = ""
)

type Endpoint struct {
	Method  string
	Path    string
	Handler fiber.Handler
}

type ServiceInfo struct {
	Name       string
	IP         string
	ApiVersion string
	Path       string
}

type Service struct {
	Info      ServiceInfo
	Endpoints []Endpoint
}

type FiberServer struct {
	App *fiber.App
}

func (s FiberServer) RegisterRoute(e *Endpoint, prefix string) {
	s.App.Add(e.Method, prefix+e.Path, e.Handler)
}

func (s FiberServer) RegisterService(service Service) {
	prefix := service.Info.ApiVersion + "/" + service.Info.Path + "/"
	for _, e := range service.Endpoints {
		s.RegisterRoute(&e, prefix)
	}
}
