package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func ForwardToTicketService(c *fiber.Ctx) error {
	addr := TicketServiceIP + c.OriginalURL()
	return proxy.Forward(addr)(c)
}

func GetTicket(c *fiber.Ctx) error {
	addr := TicketServiceIP + c.OriginalURL()
	return proxy.Forward(addr)(c)
}

func NewTicketService() Service {
	service := Service{
		Info: ServiceInfo{
			Name:       "Ticket",
			IP:         TicketServiceIP,
			ApiVersion: ApiVersion,
			Path:       "tickets",
		},
		Endpoints: []Endpoint{
			{"GET", "", ForwardToTicketService},
			{"POST", "", ForwardToTicketService},
			{"GET", ":ticketUid", GetTicket},
			{"DELETE", ":ticketUid", GetTicket},
		},
	}
	return service
}
