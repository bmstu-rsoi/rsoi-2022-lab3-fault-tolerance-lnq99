package services

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

type MeResponse struct {
	Tickets   []TicketResponse      `json:"tickets"`
	Privilege PrivilegeInfoResponse `json:"privilege,omitempty"`
}

type TicketResponseStatus string

// TicketResponse defines model for TicketResponse.
type TicketResponse struct {
	// Дата и время вылета
	Date time.Time `json:"date,omitempty"`

	// Номер полета
	FlightNumber string `json:"flightNumber,omitempty"`

	// Страна и аэропорт прибытия
	FromAirport string `json:"fromAirport,omitempty"`

	// Стоимость
	Price int32 `json:"price,omitempty"`

	// Статус билета
	Status TicketResponseStatus `json:"status,omitempty"`

	// UUID билета
	TicketUid string `json:"ticketUid,omitempty"`

	// Страна и аэропорт прибытия
	ToAirport string `json:"toAirport,omitempty"`
}

func GetMe(c *fiber.Ctx) error {
	url1 := fmt.Sprintf("%s/%s/tickets", TicketServiceIP, ApiVersion)
	url2 := fmt.Sprintf("%s/%s/privilege", BonusServiceIP, ApiVersion)
	username := c.GetReqHeaders()[UsernameHeader]
	var tickets []TicketResponse

	a := fiber.AcquireAgent()

	req := a.Request()
	req.Header.SetMethod(fiber.MethodGet)

	{
		req.SetRequestURI(url1)
		req.Header.Set(UsernameHeader, username)
		if err := a.Parse(); err != nil {
			panic(err)
		}
		code, body, errs := a.Bytes()
		log.Println(code, errs)
		json.Unmarshal(body, &tickets)
	}

	//{
	//	req.SetRequestURI(url2)
	//	req.Header.Set(UsernameHeader, username)
	//	if err := a.Parse(); err != nil {
	//		panic(err)
	//	}
	//	code, body, errs := a.Bytes()
	//	log.Println(code, errs)
	//	json.Unmarshal(body, &r.Privilege)
	//}

	var privilege PrivilegeInfoResponse

	_, err := bonusCb.Execute(func() (interface{}, error) {
		req.SetRequestURI(url2)
		req.Header.Set(UsernameHeader, username)
		if err := a.Parse(); err != nil {
			return nil, err
		}
		code, body, errs := a.Bytes()
		log.Println(code, errs)
		if len(errs) != 0 {
			return nil, errs[0]
		}
		json.Unmarshal(body, &privilege)
		return nil, nil
	})

	log.Println(err)

	return c.JSON(MeResponse{
		Tickets:   tickets,
		Privilege: privilege,
	})
}
