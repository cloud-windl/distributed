package log

import (
	"distributed/cmd/errno"
	"distributed/pkg/response"
	"io"

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
