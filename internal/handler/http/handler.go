package handler

import (
	"github.com/Dolald/smartway_test_work/configs"
	"github.com/Dolald/smartway_test_work/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
	cfg     *configs.HandlerConfig
}

func NewHandler(services *service.Service, cfg *configs.HandlerConfig) *Handler {
	return &Handler{service: services,
		cfg: cfg}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	employees := router.Group("/employees")
	{
		employees.POST("/companies/departments/add_employee", h.createEmployee)
		employees.GET("/companies/departments/:id/employees", h.getEmployeesByDepartmentId)
		employees.GET("/companies/:id/employees", h.getCompanyEmployees)
		employees.DELETE("/:id", h.deleteEmployee)
		employees.PUT("/:id", h.updateEmployee)
	}

	return router
}
