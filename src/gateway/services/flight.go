package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func GetFlights(c *fiber.Ctx) error {
	addr := FlightServiceIP + c.OriginalURL()
	return proxy.Forward(addr)(c)
}

func ForwardToFlightService(c *fiber.Ctx) error {
	addr := FlightServiceIP + c.OriginalURL()
	return proxy.Forward(addr)(c)
}

func NewFlightService() Service {
	service := Service{
		Info: ServiceInfo{
			Name:       "Flight",
			IP:         FlightServiceIP,
			ApiVersion: ApiVersion,
			Path:       "flights",
		},
		Endpoints: []Endpoint{
			{"GET", "", ForwardToFlightService},
			{"GET", ":flightNumber", ForwardToFlightService},
		},
	}
	return service
}
