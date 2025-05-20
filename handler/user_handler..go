package handler

import (
	"net/http"
	"user-service/controller"
	"user-service/model"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Ctrl *controller.UserController
}

func (h *UserHandler) Register(c *gin.Context) {
	var body struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	user := &model.User{Name: body.Name, Email: body.Email, Phone: body.Phone}
	if err := h.Ctrl.Register(user, body.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not register"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "registered"})
}

func (h *UserHandler) Login(c *gin.Context) {
	var body struct {
		Identifier string `json:"identifier"`
		Password   string `json:"password"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	token, err := h.Ctrl.Login(body.Identifier, body.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
