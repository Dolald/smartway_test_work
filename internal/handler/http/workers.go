package handler

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"
	"time"

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

	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
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
	var input models.UpdatedEmployeeRequest

	if err := c.BindJSON(&input); err != nil {
		slog.Error("BindJSON failed", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	err := h.service.Employee.UpdateEmployee(ctx, input)
	if err != nil {
		slog.Error("UpdateEmployee failed", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falied to update employee"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee updated successfully", "employee": input})
}

func (h *Handler) getWorkersCompanyDepartment(c *gin.Context) {
	id := c.Param("id")
	departmentId, _ := strconv.Atoi(id)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	employeeList, err := h.service.Employee.GetEmployeesCompanyDepartment(ctx, departmentId)
	if err != nil {
		slog.Error("GetWorkersCompanyDepartment failed", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falied to "})
		return
	}

	c.JSON(http.StatusOK, employeeList)
}

func (h *Handler) getCompanyWorkers(c *gin.Context) {

}

func (h *Handler) deleteWorker(c *gin.Context) {

}
