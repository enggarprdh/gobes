package controllers

import (
	"gobes/pkg/response"
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
