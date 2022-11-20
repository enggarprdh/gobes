package controllers

import (
	"fmt"
	"gobes/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Example struct {
	Name string `json:"name"`
}

func (e *Example) Ping(c *gin.Context) {
	response.Success(c, &response.SuccessResponse{
		Code: http.StatusOK,
		Data: "Ping success",
	})
}

func (e *Example) SayName(c *gin.Context) {

	err := c.ShouldBindJSON(&e)
	if err != nil {
		response.Fail(c, &response.FailResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	response.Success(c, &response.SuccessResponse{
		Code:    http.StatusOK,
		Message: fmt.Sprintf("Your name is %s", e.Name),
	})
}
