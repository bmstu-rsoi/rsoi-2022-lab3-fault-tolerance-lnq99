package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"ticket/model"
	"ticket/repository"

	"github.com/google/uuid"
)

const GatewayUrl = "http://gateway:8080"

type Service interface {
	ListTickets(ctx context.Context, username string) []*model.TicketResponse
	GetTicket(ctx context.Context, username, ticketUid string) *model.TicketResponse
	CreateTicket(ctx context.Context, username string, ticketReq *model.TicketPurchaseRequest) (*model.TicketPurchaseResponse, error)
	DeleteTicket(ctx context.Context, username, ticketUid string) error
}

type serviceImpl struct {
	repo repository.Repo
}

func NewService(repo repository.Repo) Service {
	return &serviceImpl{repo: repo}
}

func toTicketResponse(t *model.Ticket) *model.TicketResponse {

	url := fmt.Sprintf("%s/%s/%s", GatewayUrl, "api/v1", "flights"+"/"+t.FlightNumber)
	res, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer res.Body.Close()

	flight := model.FlightResponse{}
	err = json.NewDecoder(res.Body).Decode(&flight)
	if err != nil {
		return nil
	}

	return &model.TicketResponse{
		Date:         flight.Date,
		FlightNumber: t.FlightNumber,
		FromAirport:  flight.FromAirport,
		Price:        t.Price,
		Status:       model.TicketResponseStatus(t.Status),
		TicketUid:    t.TicketUid.String(),
		ToAirport:    flight.ToAirport,
	}
}

func (s *serviceImpl) ListTickets(ctx context.Context, username string) []*model.TicketResponse {
	tickets, err := s.repo.ListTickets(ctx, username)
	if err != nil {
		return []*model.TicketResponse{}
	}
	r := make([]*model.TicketResponse, len(tickets))
	for i, t := range tickets {
		t1 := model.Ticket(t)
		r[i] = toTicketResponse(&t1)
	}
	return r
}

func (s *serviceImpl) GetTicket(ctx context.Context, username, ticketUid string) *model.TicketResponse {
	uid, err := uuid.Parse(ticketUid)
	if err != nil {
		return nil
	}
	ticket, err := s.repo.GetTicket(ctx, repository.GetTicketParams{
		Username:  username,
		TicketUid: uid,
	})
	if err != nil {
		return nil
	}
	r := model.Ticket(ticket)
	return toTicketResponse(&r)
}

func (s *serviceImpl) CreateTicket(ctx context.Context, username string, ticketReq *model.TicketPurchaseRequest) (*model.TicketPurchaseResponse, error) {
	ticketUid, _ := uuid.NewUUID()
	t, err := s.repo.CreateTicket(ctx, repository.CreateTicketParams{
		TicketUid:    ticketUid,
		Username:     username,
		FlightNumber: ticketReq.FlightNumber,
		Price:        ticketReq.Price,
		Status:       "PAID",
	})

	var url string
	flight := model.FlightResponse{}
	privilege := model.PrivilegeInfoResponse{}
	client := &http.Client{}

	{
		url = fmt.Sprintf("%s/%s/%s", GatewayUrl, "api/v1", "flights"+"/"+t.FlightNumber)
		req, _ := http.NewRequest("GET", url, nil)
		res, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		err = json.NewDecoder(res.Body).Decode(&flight)
		if err != nil {
			return nil, err
		}
	}
	{
		url = fmt.Sprintf("%s/%s/%s", GatewayUrl, "api/v1", "privilege")
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("X-User-Name", t.Username)
		res, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		err = json.NewDecoder(res.Body).Decode(&privilege)
		if err != nil {
			return nil, err
		}
	}

	r := model.TicketPurchaseResponse{
		Date:          flight.Date,
		FlightNumber:  t.FlightNumber,
		FromAirport:   flight.FromAirport,
		PaidByBonuses: 0,
		PaidByMoney:   t.Price,
		Price:         t.Price,
		Privilege: model.PrivilegeShortInfo{
			Balance: privilege.Balance,
			Status:  model.PrivilegeShortInfoStatus(privilege.Status),
		},
		Status:    model.TicketPurchaseResponseStatus(t.Status),
		TicketUid: t.TicketUid.String(),
		ToAirport: flight.ToAirport,
	}

	balanceHistory := model.BalanceHistory{
		Date:      flight.Date,
		TicketUid: t.TicketUid,
	}

	if ticketReq.PaidFromBalance {
		r.PaidByBonuses = int32(math.Max(float64(privilege.Balance), float64(t.Price)))
		r.PaidByMoney = t.Price - r.PaidByBonuses
		r.Privilege.Balance = r.Privilege.Balance - r.PaidByBonuses
		balanceHistory.BalanceDiff = -r.PaidByBonuses
		balanceHistory.OperationType = "DEBIT_THE_ACCOUNT"
	} else {
		r.Privilege.Balance = r.Privilege.Balance + t.Price/10
		balanceHistory.BalanceDiff = t.Price / 10
		balanceHistory.OperationType = "FILL_IN_BALANCE"
	}

	{
		url = fmt.Sprintf("%s/%s/%s", GatewayUrl, "api/v1", "privilege")
		body, _ := json.Marshal(balanceHistory)

		req, _ := http.NewRequest("POST", url, bytes.NewReader(body))
		req.Header.Set("X-User-Name", t.Username)
		res, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()
	}

	return &r, err
}

func (s *serviceImpl) DeleteTicket(ctx context.Context, username, ticketUid string) error {
	uid, err := uuid.Parse(ticketUid)
	if err != nil {
		return err
	}
	err = s.repo.DeleteTicket(ctx, repository.DeleteTicketParams{
		Username:  username,
		TicketUid: uid,
	})
	return err
}
