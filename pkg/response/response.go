package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SuccessStatus = "success"
	FailStatus    = "fail"
)

type SuccessResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type FailResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"errors"`
}

func Success(c *gin.Context, r *SuccessResponse) {
	r.Status = SuccessStatus
	if r.Code == 0 {
		r.Code = http.StatusOK
	}

	if r.Message == "" {
		r.Message = SuccessStatus
	}

	c.JSON(r.Code, r)
}

func Fail(c *gin.Context, r *FailResponse) {
	r.Status = FailStatus
	if r.Code == 0 {
		r.Code = http.StatusOK
	}

	if r.Message == "" {
		r.Message = FailStatus
	}

	c.JSON(r.Code, r)
}
