package handler

import (
	"log/slog"

	"github.com/Dolald/smartway_test_work/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
	logger  *slog.Logger
}

func NewHandler(services *service.Service, logger *slog.Logger) *Handler {
	return &Handler{service: services,
		logger: logger,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	documents := router.Group("/workers")
	{
		documents.POST("/", h.createWorker)
		documents.GET("/:id", h.getWorkersCompanyDepartment) // получение всех сотрудников отдела компании
		documents.GET("/:id", h.getCompanyWorkers)           // получение всех сотрудников компании                                                 // получение сотрудника по id
		documents.DELETE("/:id", h.deleteWorker)
		documents.PUT("/:id", h.updateWorker)
	}

	return router
}
