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
		slog.Error("bindJSON failed", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input data"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), configs.ContextTime)
	defer cancel()

	workerId, err := h.service.Employee.CreateEmployee(ctx, input)
	if err != nil {
		slog.Error("createEmployee failed", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falied to create employee"})
		return
	}

	c.JSON(http.StatusOK, workerId)
}

func (h *Handler) updateEmployee(c *gin.Context) {
	id := c.Param(configs.Id)
	employeetId, err := strconv.Atoi(id)
	if err != nil {
		slog.Error("get department id failed", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "falied to get department id"})
		return
	}

	var input models.UpdateEmployeeRequest

	if err := c.BindJSON(&input); err != nil {
		slog.Error("bindJSON failed", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), configs.ContextTime)
	defer cancel()

	err = h.service.Employee.UpdateEmployee(ctx, input, employeetId)
	if err != nil {
		slog.Error("updateEmployee failed", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falied to update employee"})
		return
	}

	c.JSON(http.StatusOK, input)
}

func (h *Handler) getEmployeesCompanyDepartment(c *gin.Context) {
	id := c.Param(configs.Id)
	departmentId, err := strconv.Atoi(id)
	if err != nil {
		slog.Error("get department id failed", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "falied to get department id"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), configs.ContextTime)
	defer cancel()

	employeesList, err := h.service.Employee.GetEmployeesCompanyDepartment(ctx, departmentId)
	if err != nil {
		slog.Error("getEmployeesCompanyDepartment failed", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falied to get employee list"})
		return
	}

	c.JSON(http.StatusOK, employeesList)
}

func (h *Handler) getCompanyEmployees(c *gin.Context) {
	id := c.Param(configs.Id)
	companyId, err := strconv.Atoi(id)
	if err != nil {
		slog.Error("get company id failed", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "falied to get company id"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), configs.ContextTime)
	defer cancel()

	employeesList, err := h.service.Employee.GetEmployeesCompany(ctx, companyId)
	if err != nil {
		slog.Error("getEmployeesCompany failed", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falied to get employee list"})
		return
	}

	c.JSON(http.StatusOK, employeesList)
}

func (h *Handler) deleteEmployee(c *gin.Context) {
	id := c.Param(configs.Id)
	employeeId, err := strconv.Atoi(id)
	if err != nil {
		slog.Error("get employee id failed", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "falied to get employee id"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), configs.ContextTime)
	defer cancel()

	err = h.service.Employee.DeleteEmployee(ctx, employeeId)
	if err != nil {
		slog.Error("deleteEmployee failed", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falied to delete employee"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}
