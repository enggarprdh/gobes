package validator

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type RequestForm interface {
	Message() govalidator.MapData
	Rules() govalidator.MapData
}

type NoMessageForm interface {
	Rules() govalidator.MapData
}

func Validate(c *gin.Context, form RequestForm) error {
	fmt.Println(c.Request)
	opt := govalidator.Options{
		Request:  c.Request,
		Rules:    form.Rules(),
		Messages: form.Message(),
	}
	res := govalidator.New(opt).Validate()

	if len(res) > 0 {
		for _, x := range res {
			return errors.New(x[0])
		}
	}
	return nil
}

func DefaultMessageValidate(c *gin.Context, form NoMessageForm) url.Values {
	opt := govalidator.Options{
		Request: c.Request,
		Rules:   form.Rules(),
	}
	return govalidator.New(opt).Validate()
}
