package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/wregulski/log-poc/pkg/log"
	"github.com/wregulski/log-poc/pkg/service"
)

type Handler struct {
	services *service.Services
	logger   log.Logger
}

func NewHandler(services *service.Services, lf log.LoggerFactory) *Handler {
	return &Handler{
		services: services,
		logger:   lf.NewLogger("handler"),
	}
}

func (h *Handler) Init() *gin.Engine {
	r := gin.Default()

	r.GET("/echo", h.echo)

	return r
}

func (h *Handler) echo(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": h.services.Echo.Echo(),
	})
}
