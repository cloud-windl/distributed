package grades

import (
	myjwt "distributed/pkg/jwt"
	"distributed/pkg/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthHandler struct {
	authService *AuthService
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{
		authService: NewAuthService(db),
	}
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 7001, err.Error())
		return
	}

	user, err := h.authService.Login(req.Username, req.Password)
	if err != nil {
		response.Error(c, 7002, "invalid username or password")
		return
	}

	token, err := myjwt.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		response.Error(c, 7003, "failed to generate token")
		return
	}

	response.OK(c, gin.H{
		"token": token,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"role":     user.Role,
		},
	})
}
