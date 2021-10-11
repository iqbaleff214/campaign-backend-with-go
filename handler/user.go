package handler

import (
	"bwacampaign/helper"
	"bwacampaign/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userHandler struct {
	service user.Service
}

func NewUserHandler(service user.Service) *userHandler {
	return &userHandler{service: service}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var request user.RegisterUserRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, helper.ApiResponse("Request has been rejected", http.StatusUnprocessableEntity, "error", helper.ErrorValidationResponse(err)))
		return
	}

	newUser, err := h.service.RegisterUser(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiResponse("Account already registered", http.StatusBadRequest, "error", nil))
		return
	}

	formatted := user.FormatUser(newUser, "abc")
	response := helper.ApiResponse("Account has been registered", http.StatusOK, "success", formatted)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context)  {
	var request user.LoginRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, helper.ApiResponse("Login failed", http.StatusUnprocessableEntity, "error", helper.ErrorValidationResponse(err)))
		return
	}

	loginUser, err := h.service.Login(request)
	if err != nil {
		c.JSON(http.StatusNotFound, helper.ApiResponse("Login failed", http.StatusNotFound, "error", gin.H{"error": err.Error()}))
		return
	}

	formatted := user.FormatUser(loginUser, "abc")
	response := helper.ApiResponse("Successfully logged in", http.StatusOK, "success", formatted)
	c.JSON(http.StatusOK, response)

}
