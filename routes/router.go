package rotutes

import (
	controllers "gobes/controllers"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) *gin.Engine {
	v1 := r.Group("api/v1")
	{
		//For example we create "api/v1/ping"
		ex := new(controllers.Example)
		v1.GET("/ping", ex.Ping)
	}

	return r
}
