package controller

import (
	"log"
	"net/http"

	"bonus/service"

	"github.com/lnq99/rsoi-2022-lab3-fault-tolerance-lnq99/src/pkg/model"

	"github.com/gin-gonic/gin"
)

type GinController struct {
	service service.Service
}

func NewGinController(service service.Service) *GinController {
	return &GinController{service}
}

func (c *GinController) ListPrivilegeHistories(ctx *gin.Context) {
	username := ctx.GetHeader("X-User-Name")
	r := c.service.GetPrivilege(ctx, username)
	log.Println(r)
	ctx.JSON(http.StatusOK, r)
}

func (c *GinController) UpdateBalanceAndHistory(ctx *gin.Context) {
	username := ctx.GetHeader("X-User-Name")

	history := model.BalanceHistory{}
	if err := ctx.ShouldBindJSON(&history); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	err := c.service.UpdateBalanceAndHistory(ctx, username, history)
	if err != nil {
		ctx.Status(http.StatusOK)
	} else {
		ctx.Status(http.StatusInternalServerError)
	}
}
