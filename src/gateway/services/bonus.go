package services

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/google/uuid"
)

func ForwardToBonusService(c *fiber.Ctx) error {
	addr := BonusServiceIP + c.OriginalURL()
	return proxy.Forward(addr)(c)
}

func GetBonus(c *fiber.Ctx) error {
	url := BonusServiceIP + c.OriginalURL()
	status := http.StatusOK
	var r PrivilegeInfoResponse

	_, err := bonusCb.Execute(func() (interface{}, error) {
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set(UsernameHeader, c.GetReqHeaders()[UsernameHeader])
		res, err := Client.Do(req)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		status = res.StatusCode
		err = json.NewDecoder(res.Body).Decode(&r)
		log.Println(status, r)
		if err != nil {
			return nil, err
		}

		return r, nil
	})

	if err != nil {
		log.Println("bonus get", err)
		return c.Status(http.StatusServiceUnavailable).JSON(map[string]string{
			"message": "Bonus Service unavailable",
		})
		//return c.Status(http.StatusInternalServerError).Send(nil)
	}

	return c.Status(status).JSON(r)
}

func NewBonusService() Service {
	s := Service{
		Info: ServiceInfo{
			Name:       "Bonus",
			IP:         BonusServiceIP,
			ApiVersion: ApiVersion,
			Path:       "privilege",
		},
	}

	s.Endpoints = []Endpoint{
		{"GET", "", GetBonus},
		{"POST", "", ForwardToBonusService},
	}

	return s
}

// Тип операции
type BalanceHistoryOperationType string

// Defines values for BalanceHistoryOperationType.
const (
	DEBITTHEACCOUNT BalanceHistoryOperationType = "DEBIT_THE_ACCOUNT"
	FILLEDBYMONEY   BalanceHistoryOperationType = "FILLED_BY_MONEY"
	FILLINBALANCE   BalanceHistoryOperationType = "FILL_IN_BALANCE"
)

// BalanceHistory defines model for BalanceHistory.
type BalanceHistory struct {
	// Изменение баланса
	BalanceDiff int32 `json:"balanceDiff,omitempty"`

	// Дата и время операции
	Date time.Time `json:"date,omitempty"`

	// Тип операции
	OperationType BalanceHistoryOperationType `json:"operationType,omitempty"`

	// UUID билета по которому была операция с бонусами
	TicketUid uuid.UUID `json:"ticketUid,omitempty"`
}

// PrivilegeInfoResponse defines model for PrivilegeInfoResponse.
type PrivilegeInfoResponse struct {
	// Баланс бонусного счета
	Balance int32 `json:"balance,omitempty"`

	// История изменения баланса
	History []BalanceHistory `json:"history,omitempty"`

	// Статус в бонусной программе
	Status PrivilegeInfoResponseStatus `json:"status,omitempty"`
}

// Статус в бонусной программе
type PrivilegeInfoResponseStatus string

// Defines values for PrivilegeInfoResponseStatus.
const (
	PrivilegeInfoResponseStatusBRONZE PrivilegeInfoResponseStatus = "BRONZE"
	PrivilegeInfoResponseStatusGOLD   PrivilegeInfoResponseStatus = "GOLD"
	PrivilegeInfoResponseStatusSILVER PrivilegeInfoResponseStatus = "SILVER"
)
