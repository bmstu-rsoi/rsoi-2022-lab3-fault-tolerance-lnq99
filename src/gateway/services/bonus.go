package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/lnq99/rsoi-2022-lab3-fault-tolerance-lnq99/src/pkg/model"
)

func NewBonusService() Service {
	s := Service{
		Info: ServiceInfo{
			Name:       "Bonus",
			IP:         BonusServiceIP,
			ApiVersion: ApiVersion,
			Path:       "privilege",
		},
		Endpoints: []Endpoint{
			{"GET", "", GetBonus},
			{"POST", "", ForwardToBonusService},
		},
	}
	return s
}

func ForwardToBonusService(c *fiber.Ctx) error {
	addr := BonusServiceIP + c.OriginalURL()
	return proxy.Forward(addr)(c)
}

func GetBonus(c *fiber.Ctx) error {
	url := BonusServiceIP + c.OriginalURL()
	header := map[string]string{UsernameHeader: c.GetReqHeaders()[UsernameHeader]}

	r, err := CallServiceWithCircuitBreaker(
		bonusCb, "GET", url, header, nil)

	return fiberProcessResponse[model.PrivilegeInfoResponse](c, r.status, r.body, err)
}
