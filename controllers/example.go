package controllers

import (
	"fmt"
	"gobes/models"
	"gobes/pkg/response"
	"gobes/pkg/utils/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Example struct {
}

func (e *Example) Ping(c *gin.Context) {
	response.Success(c, &response.SuccessResponse{
		Code: http.StatusOK,
		Data: "Ping success",
	})
}

func (e *Example) SayName(c *gin.Context) {
	var i models.ExampleParam
	err := validator.Validate(c, &i)
	if err != nil {
		response.Fail(c, &response.FailResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	response.Success(c, &response.SuccessResponse{
		Code:    http.StatusOK,
		Message: fmt.Sprintf("Your name is %s", i.Name),
	})
}
