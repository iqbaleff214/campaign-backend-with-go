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
		c.JSON(http.StatusUnprocessableEntity, helper.ApiResponse("Request has been rejected", http.StatusUnprocessableEntity, "error", helper.ErrorResponse(err)))
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
