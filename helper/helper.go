package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    uint16  `json:"code"`
	Status  string `json:"status"`
}

func ApiResponse(message string, code uint16, status string, data interface{}) Response {
	metaJson := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}
	responseJson := Response{
		Meta: metaJson,
		Data: data,
	}

	return responseJson
}

func ErrorValidationResponse(err error) gin.H {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}
	return gin.H{"errors": errors}
}
