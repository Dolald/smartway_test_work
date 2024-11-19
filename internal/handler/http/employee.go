package handler

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Dolald/smartway_test_work/configs"
	"github.com/Dolald/smartway_test_work/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createEmployee(c *gin.Context) {
	var input models.EmployeeRequest

	if err := c.BindJSON(&input); err != nil {
		slog.Error("BindJSON failed", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), configs.ContextTime)
	defer cancel()

	workerId, err := h.service.Employee.CreateEmployee(ctx, input)
	if err != nil {
		slog.Error("CreateEmployee failed", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falied to create employee"})
		return
	}

	c.JSON(http.StatusOK, workerId)
}

func (h *Handler) updateWorker(c *gin.Context) {
	var input models.UpdateEmployeeRequest

	if err := c.BindJSON(&input); err != nil {
		slog.Error("BindJSON failed", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), configs.ContextTime)
	defer cancel()

	err := h.service.Employee.UpdateEmployee(ctx, input)
	if err != nil {
		slog.Error("UpdateEmployee failed", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falied to update employee"})
		return
	}

	c.JSON(http.StatusOK, input)
}

func (h *Handler) getEmployeesCompanyDepartment(c *gin.Context) {
	id := c.Param("id")
	departmentId, _ := strconv.Atoi(id)

	ctx, cancel := context.WithTimeout(c.Request.Context(), configs.ContextTime)
	defer cancel()

	employeesList, err := h.service.Employee.GetEmployeesCompanyDepartment(ctx, departmentId)
	if err != nil {
		slog.Error("GetEmployeesCompanyDepartment failed", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falied to get department list"})
		return
	}

	c.JSON(http.StatusOK, employeesList)
}

func (h *Handler) getCompanyEmployees(c *gin.Context) {
	id := c.Param("id")
	companyId, _ := strconv.Atoi(id)

	ctx, cancel := context.WithTimeout(c.Request.Context(), configs.ContextTime)
	defer cancel()

	employeesList, err := h.service.Employee.GetEmployeesCompany(ctx, companyId)
	if err != nil {
		slog.Error("GetEmployeesCompanyDepartment failed", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falied to get department list"})
		return
	}

	c.JSON(http.StatusOK, employeesList)
}

func (h *Handler) deleteEmployee(c *gin.Context) {
	id := c.Param("id")
	employeeId, _ := strconv.Atoi(id)

	ctx, cancel := context.WithTimeout(c.Request.Context(), configs.ContextTime)
	defer cancel()

	err := h.service.Employee.DeleteEmployee(ctx, employeeId)
	if err != nil {
		slog.Error("GetEmployeesCompanyDepartment failed", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falied to get department list"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}
