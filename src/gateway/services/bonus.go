package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func ForwardToBonusService(c *fiber.Ctx) error {
	addr := BonusServiceIP + c.OriginalURL()
	return proxy.Forward(addr)(c)
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
		{"GET", "", ForwardToBonusService},
		{"POST", "", ForwardToBonusService},
	}

	return s
}
