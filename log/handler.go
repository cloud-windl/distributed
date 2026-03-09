package log

import (
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
	if err != nil || len(msg) == 0 {
		c.Status(http.StatusBadRequest)
		return
	}

	h.service.Write(string(msg))
	c.Status(http.StatusOK)
}
