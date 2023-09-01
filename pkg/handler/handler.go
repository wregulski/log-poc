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
	h.logger.Info("Initializing handler...")

	h.logger.Warn("Reminder to add some middleware here! (Just testing the logger)")

	r.GET("/echo", h.echo)
	r.GET("/error-echo", h.errorEcho)

	return r
}

func (h *Handler) echo(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": h.services.Echo.Echo(),
	})
}

func (h *Handler) errorEcho(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": h.services.Echo.ErrorEcho(),
	})
}
