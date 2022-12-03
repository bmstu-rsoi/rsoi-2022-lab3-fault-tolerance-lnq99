package model

import "time"

type Airport struct {
	ID      int32  `json:"id"`
	Name    string `json:"name"`
	City    string `json:"city"`
	Country string `json:"country"`
}

type Flight struct {
	ID            int32     `json:"id"`
	FlightNumber  string    `json:"flightNumber"`
	Datetime      time.Time `json:"datetime"`
	FromAirportID int32     `json:"fromAirportID"`
	ToAirportID   int32     `json:"toAirportID"`
	Price         int32     `json:"price"`
}

type FlightResponse struct {
	FlightNumber string    `json:"flightNumber,omitempty"` // Номер полета
	FromAirport  string    `json:"fromAirport,omitempty"`  // Страна и аэропорт прибытия
	ToAirport    string    `json:"toAirport,omitempty"`    // Страна и аэропорт прибытия
	Date         time.Time `json:"date,omitempty"`         // Дата и время вылета
	Price        int32     `json:"price,omitempty"`        // Стоимость
}

type PaginationResponse struct {
	TotalElements int32            `json:"totalElements,omitempty"` // Общее количество элементов
	Page          int32            `json:"page,omitempty"`          // Номер страницы
	PageSize      int32            `json:"pageSize,omitempty"`      // Количество элементов на странице
	Items         []FlightResponse `json:"items,omitempty"`
}
