package model

import (
	"time"

	"github.com/google/uuid"
)

type Ticket struct {
	ID           int32     `json:"id"`
	TicketUid    uuid.UUID `json:"ticketUid"`
	Username     string    `json:"username"`
	FlightNumber string    `json:"flightNumber"`
	Price        int32     `json:"price"`
	Status       string    `json:"status"`
}

type FlightResponse struct {
	FlightNumber string    `json:"flightNumber,omitempty"` // Номер полета
	FromAirport  string    `json:"fromAirport,omitempty"`  // Страна и аэропорт прибытия
	ToAirport    string    `json:"toAirport,omitempty"`    // Страна и аэропорт прибытия
	Date         time.Time `json:"date,omitempty"`         // Дата и время вылета
	Price        int32     `json:"price,omitempty"`        // Стоимость
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

// Статус в бонусной программе
type PrivilegeInfoResponseStatus string

type PrivilegeInfoResponse struct {
	// Баланс бонусного счета
	Balance int32 `json:"balance"`

	// История изменения баланса
	History []BalanceHistory `json:"history"`

	// Статус в бонусной программе
	Status PrivilegeInfoResponseStatus `json:"status"`
}

// Статус билета
type TicketPurchaseResponseStatus string
type TicketResponseStatus string

// Defines values for TicketPurchaseResponseStatus.
const (
	TicketPurchaseResponseStatusCANCELED TicketPurchaseResponseStatus = "CANCELED"
	TicketPurchaseResponseStatusPAID     TicketPurchaseResponseStatus = "PAID"
)

// Defines values for TicketResponseStatus.
const (
	TicketResponseStatusCANCELED TicketResponseStatus = "CANCELED"
	TicketResponseStatusPAID     TicketResponseStatus = "PAID"
)

// PrivilegeShortInfo defines model for PrivilegeShortInfo.
type PrivilegeShortInfo struct {
	// Баланс бонусного счета
	Balance int32 `json:"balance"`

	// Статус в бонусной программе
	Status PrivilegeShortInfoStatus `json:"status"`
}

// Статус в бонусной программе
type PrivilegeShortInfoStatus string

// TicketPurchaseRequest defines model for TicketPurchaseRequest.
type TicketPurchaseRequest struct {
	// Номер полета
	FlightNumber string `json:"flightNumber,omitempty"`

	// Выполнить списание бонусных баллов в учет покупки билета
	PaidFromBalance bool `json:"paidFromBalance,omitempty"`

	// Стоимость
	Price int32 `json:"price,omitempty"`
}

// TicketPurchaseResponse defines model for TicketPurchaseResponse.
type TicketPurchaseResponse struct {
	// Время вылета
	Date time.Time `json:"date"`

	// Номер полета
	FlightNumber string `json:"flightNumber"`

	// Страна и аэропорт прибытия
	FromAirport string `json:"fromAirport"`

	// Сумма оплаченная бонусами
	PaidByBonuses int32 `json:"paidByBonuses"`

	// Сумма оплаченная деньгами
	PaidByMoney int32 `json:"paidByMoney"`

	// Стоимость
	Price     int32              `json:"price"`
	Privilege PrivilegeShortInfo `json:"privilege"`

	// Статус билета
	Status TicketPurchaseResponseStatus `json:"status"`

	// UUID билета
	TicketUid string `json:"ticketUid"`

	// Страна и аэропорт прибытия
	ToAirport string `json:"toAirport"`
}

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

// ErrorDescription defines model for ErrorDescription.
type ErrorDescription struct {
	Error *string `json:"error,omitempty"`
	Field *string `json:"field,omitempty"`
}

// ValidationErrorResponse defines model for ValidationErrorResponse.
type ValidationErrorResponse struct {
	// Массив полей с описанием ошибки
	Errors *[]ErrorDescription `json:"errors,omitempty"`

	// Информация об ошибке
	Message *string `json:"message,omitempty"`
}

// GetTicketsParams defines parameters for GetTickets.
type GetTicketsParams struct {
	// Имя пользователя
	XUserName string `json:"X-User-Name"`
}

// PostTicketsJSONBody defines parameters for PostTickets.
type PostTicketsJSONBody = TicketPurchaseRequest

// PostTicketsParams defines parameters for PostTickets.
type PostTicketsParams struct {
	// Имя пользователя
	XUserName string `json:"X-User-Name"`
}

// DeleteTicketsTicketUidParams defines parameters for DeleteTicketsTicketUid.
type DeleteTicketsTicketUidParams struct {
	// Имя пользователя
	XUserName string `json:"X-User-Name"`
}

// GetTicketsTicketUidParams defines parameters for GetTicketsTicketUid.
type GetTicketsTicketUidParams struct {
	// Имя пользователя
	XUserName string `json:"X-User-Name"`
}

// PostTicketsJSONRequestBody defines body for PostTickets for application/json ContentType.
type PostTicketsJSONRequestBody = PostTicketsJSONBody
