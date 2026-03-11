package registry

import (
	"distributed/cmd/errno"
	"distributed/pkg/response"
	"strings"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler() *Handler {
	return &Handler{
		service: NewService(),
	}
}

func (h *Handler) RegisterService(c *gin.Context) {
	var req Registration
	if err := c.ShouldBind(&req); err != nil {
		response.Error(c, errno.CodeInvalidRegistryReq, "invalid registration request")
		return
	}
	err := h.service.Add(req)
	if err != nil {
		response.Error(c, errno.CodeRegisterFailed, err.Error())
		return
	}
	response.OK(c, gin.H{
		"message": "service registered successfully",
	})
}

func (h *Handler) DeregisterService(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		response.Error(c, errno.CodeReadBodyFailed, "failed to read request body")
		return
	}
	url := strings.TrimSpace(string(body))
	if url == "" {
		response.Error(c, errno.CodeEmptyServiceURL, "empty service url")
		return
	}
	if err := h.service.Remove(url); err != nil {
		response.Error(c, errno.CodeDeregisterFailed, err.Error())
		return
	}

	response.OK(c, gin.H{
		"message": "service deregistered successfully",
	})
}
