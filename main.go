package main

import (
	"fmt"
	"gobes/pkg/config"
	routes "gobes/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := config.Load("app.yml"); err != nil {
		fmt.Println(err.Error())
		return
	}

	gin.SetMode(config.Get().Mode)
	router := gin.Default()
	routes.Init(router)

	router.Run(config.Get().Addr)
}
