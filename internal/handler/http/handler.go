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

	employees := router.Group("/employees")
	{
		employees.POST("/", h.createEmployee)
		employees.GET("/department/:id", h.getEmployeesCompanyDepartment)
		employees.GET("/company/:id", h.getCompanyEmployees)
		employees.DELETE("/:id", h.deleteEmployee)
		employees.PUT("/", h.updateWorker)
	}

	return router
}
