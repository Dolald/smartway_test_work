package handler

import (
	"github.com/Dolald/smartway_test_work/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createWorker(c *gin.Context) {
	var input models.Worker

	if err := c.BindJSON(&input); err != nil { // парсим input, где находятся созданный список заданий
		h.logger.Error("CreateDocument failed: %w", err)
		return
	}

	// ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	// defer cancel()

	// workerId, err := h.service.Document.CreateDocument(ctx, userId, document)
	// if err != nil {
	// 	h.logger.Error("CreateDocument failed: %w", err)
	// 	return
	// }

}

func (h *Handler) updateWorker(c *gin.Context) {

}

func (h *Handler) getWorkersCompanyDepartment(c *gin.Context) {

}

func (h *Handler) getCompanyWorkers(c *gin.Context) {

}

func (h *Handler) deleteWorker(c *gin.Context) {

}
