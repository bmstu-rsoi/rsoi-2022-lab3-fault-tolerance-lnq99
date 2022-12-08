package services

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func ForwardToTicketService(c *fiber.Ctx) error {
	addr := TicketServiceIP + c.OriginalURL()
	return proxy.Forward(addr)(c)
}

func GetTickets(c *fiber.Ctx) error {
	url := TicketServiceIP + c.OriginalURL()
	status := http.StatusOK
	var r []TicketResponse

	_, err := ticketCb.Execute(func() (interface{}, error) {
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set(UsernameHeader, c.GetReqHeaders()[UsernameHeader])
		res, err := Client.Do(req)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		status = res.StatusCode
		err = json.NewDecoder(res.Body).Decode(&r)
		//log.Println(status, r)
		if err != nil {
			return nil, err
		}

		return r, nil
	})

	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).Send(nil)
	}

	return c.Status(status).JSON(r)
}

func PostTicket(c *fiber.Ctx) error {
	return ForwardToTicketService(c)
}

func GetTicket(c *fiber.Ctx) error {
	return ForwardToTicketService(c)
}

func DeleteTicket(c *fiber.Ctx) error {
	return ForwardToTicketService(c)
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
			{"GET", "", GetTickets},
			{"POST", "", PostTicket},
			{"GET", ":ticketUid", GetTicket},
			{"DELETE", ":ticketUid", DeleteTicket},
		},
	}
	return service
}
