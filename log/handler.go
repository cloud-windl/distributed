package log

import (
	"distributed/pkg/errno"
	"distributed/pkg/response"
	"distributed/registry"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler() *Handler {
	return &Handler{service: NewService()}
}

func (h *Handler) WriteLog(c *gin.Context) {
	msg, err := io.ReadAll(c.Request.Body)
	if err != nil {
		response.Error(c, errno.CodeLogReadBodyFailed, "failed to read request body")
		return
	}
	if len(msg) == 0 {
		response.Error(c, errno.CodeEmptyLogMessage, "empty log message")
		return
	}

	h.service.Write(string(msg))
	response.OK(c, gin.H{
		"message": "log written successfully",
	})
}

func (h *Handler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (h *Handler) HandleRegistryUpdate(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 9101,
			"msg":  "failed to read registry update body",
		})
		return
	}

	if err := registry.ApplyPatchBody(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 9102,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "registry update received",
	})
}
