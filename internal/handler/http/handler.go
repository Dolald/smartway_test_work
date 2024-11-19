package handler

import (
	"github.com/Dolald/smartway_test_work/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{service: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	documents := router.Group("/workers")
	{
		documents.POST("/", h.createEmployee)
		documents.GET("/department/:id", h.getWorkersCompanyDepartment)
		documents.GET("/company/:id", h.getCompanyWorkers)
		documents.DELETE("/:id", h.deleteWorker)
		documents.PUT("/", h.updateWorker)
	}

	return router
}
