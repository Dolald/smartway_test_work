package handler

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Dolald/smartway_test_work/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createEmployee(c *gin.Context) {
	var input models.CreateEmployeeRequest

	if err := c.BindJSON(&input); err != nil {
		slog.Error("bindJSON failed", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input data"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), h.cfg.RequestTimeout)
	defer cancel()

	employeeId, err := h.service.Employee.CreateEmployee(ctx, input)
	if err != nil {
		slog.Error("CreateEmployee failed", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falied to create employee"})
		return
	}

	c.JSON(http.StatusOK, employeeId)
}

func (h *Handler) updateEmployee(c *gin.Context) {
	id := c.Param(h.cfg.UrlId)
	employeetId, err := strconv.Atoi(id)
	if err != nil {
		slog.Error("get employee id failed", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "falied to get employee id"})
		return
	}

	var input models.UpdateEmployeeRequest

	if err := c.BindJSON(&input); err != nil {
		slog.Error("bindJSON failed", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), h.cfg.RequestTimeout)
	defer cancel()

	err = h.service.Employee.UpdateEmployee(ctx, input, employeetId)
	if err != nil {
		slog.Error("UpdateEmployee failed", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falied to update employee"})
		return
	}

	c.JSON(http.StatusOK, input)
}

func (h *Handler) getEmployeesByDepartmentId(c *gin.Context) {
	id := c.Param(h.cfg.UrlId)
	departmentId, err := strconv.Atoi(id)
	if err != nil {
		slog.Error("get department id failed", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "falied to get department id"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), h.cfg.RequestTimeout)
	defer cancel()

	employeesList, err := h.service.Employee.GetEmployeesByDepartmentId(ctx, departmentId)
	if err != nil {
		slog.Error("GetEmployeesByDepartmentId failed", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falied to get employee list"})
		return
	}

	c.JSON(http.StatusOK, employeesList)
}

func (h *Handler) getEmployeesByCompanyId(c *gin.Context) {
	id := c.Param(h.cfg.UrlId)
	companyId, err := strconv.Atoi(id)
	if err != nil {
		slog.Error("get company id failed", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "falied to get company id"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), h.cfg.RequestTimeout)
	defer cancel()

	employeesList, err := h.service.Employee.GetEmployeesByCompanyId(ctx, companyId)
	if err != nil {
		slog.Error("GetEmployeesByCompanyId failed", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falied to get employee list"})
		return
	}

	c.JSON(http.StatusOK, employeesList)
}

func (h *Handler) deleteEmployee(c *gin.Context) {
	id := c.Param(h.cfg.UrlId)
	employeeId, err := strconv.Atoi(id)
	if err != nil {
		slog.Error("get employee id failed", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "falied to get employee id"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), h.cfg.RequestTimeout)
	defer cancel()

	err = h.service.Employee.DeleteEmployee(ctx, employeeId)
	if err != nil {
		slog.Error("DeleteEmployee failed", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falied to delete employee"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}
