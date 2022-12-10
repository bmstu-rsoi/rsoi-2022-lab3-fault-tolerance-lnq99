package server

import (
	"bonus/config"
	"bonus/controller"
	"bonus/service"
	"net/http"

	"github.com/lnq99/rsoi-2022-lab3-fault-tolerance-lnq99/src/pkg/server"

	limits "github.com/gin-contrib/size"
	"github.com/gin-gonic/gin"
)

type GinServer struct {
	server.BaseServer
	engine   *gin.Engine
	handlers *controller.GinController
	config   *config.ServerConfig
}

func NewGinServer(service service.Service, cfg *config.ServerConfig) server.Server {
	// engine.SetMode(engine.ReleaseMode)
	engine := gin.New()

	// Middleware
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.Use(limits.RequestSizeLimiter(1000))

	ctrl := controller.NewGinController(service)
	s := GinServer{
		engine:   engine,
		handlers: ctrl,
		config:   cfg,
	}
	s.SetupRouter()
	return &s
}

func (s *GinServer) SetupRouter() {
	r := s.engine
	v1 := r.Group("api/v1")
	{
		v1.GET("privilege", s.handlers.ListPrivilegeHistories)
		v1.POST("privilege", s.handlers.UpdateBalanceAndHistory)
	}
	r.GET("/manage/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
}

func (s *GinServer) Run() {
	cfg := server.DefaultConfig()
	cfg.Addr = s.config.Host + ":" + s.config.Port
	cfg.Handler = s.engine
	s.InitHttpServer(cfg)
	s.BaseServer.Run()
}
