package model

import (
	"time"

	"github.com/google/uuid"
)

type Privilege struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	Status   string `json:"status"`
	Balance  int32  `json:"balance"`
}

type PrivilegeHistory struct {
	ID            int32     `json:"id"`
	PrivilegeID   int32     `json:"privilegeID"`
	TicketUid     uuid.UUID `json:"ticketUid"`
	Datetime      time.Time `json:"datetime"`
	BalanceDiff   int32     `json:"balanceDiff"`
	OperationType string    `json:"operationType"`
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
	Balance int32 `json:"balance"`

	// История изменения баланса
	History []BalanceHistory `json:"history"`

	// Статус в бонусной программе
	Status PrivilegeInfoResponseStatus `json:"status"`
}

// Статус в бонусной программе
type PrivilegeInfoResponseStatus string

// Defines values for PrivilegeInfoResponseStatus.
const (
	PrivilegeInfoResponseStatusBRONZE PrivilegeInfoResponseStatus = "BRONZE"
	PrivilegeInfoResponseStatusGOLD   PrivilegeInfoResponseStatus = "GOLD"
	PrivilegeInfoResponseStatusSILVER PrivilegeInfoResponseStatus = "SILVER"
)
