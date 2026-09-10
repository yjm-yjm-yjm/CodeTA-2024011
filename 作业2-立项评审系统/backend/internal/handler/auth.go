package handler

import (
	"crypto/sha1"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"hackathon-project-review/internal/database"
	"hackathon-project-review/internal/model"
	"hackathon-project-review/internal/pkg/jwt"
)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler { return &AuthHandler{} }

type registerReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"required,oneof=student reviewer"`
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req registerReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}
	sum := sha1.Sum([]byte(req.Password))
	user := model.User{Email: req.Email, PasswordHash: fmt.Sprintf("%x", sum), Role: model.Role(req.Role)}
	if err := database.DB.Create(&user).Error; err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			c.JSON(http.StatusConflict, gin.H{"error": "邮箱已注册"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注册失败"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": user.ID, "email": user.Email, "role": user.Role})
}

type loginReq struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req loginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}
	var user model.User
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "邮箱或密码错误"})
		return
	}
	sum := sha1.Sum([]byte(req.Password))
	if user.PasswordHash != fmt.Sprintf("%x", sum) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "邮箱或密码错误"})
		return
	}
	token, err := jwt.Generate(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "登录失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token, "role": user.Role})
}

func (h *AuthHandler) Me(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, user)
}
