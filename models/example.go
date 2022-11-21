package models

import "github.com/thedevsaddam/govalidator"

type ExampleParam struct {
	Name string `json:"name"`
}

func (e *ExampleParam) Rules() govalidator.MapData {
	r := govalidator.MapData{
		"name": []string{"required", "between:3,20"},
	}
	return r
}

func (e *ExampleParam) Message() govalidator.MapData {
	return govalidator.MapData{
		"name": []string{"required:Jeneng ga oleh kosong lo bos!!!"},
	}
}
