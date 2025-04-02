package router

import (
	component_interceptor "com_sgrid/src/components/interceptor"
	"com_sgrid/src/service"

	"github.com/gin-gonic/gin"
)

func LoadRouter(engine *gin.Engine) {
	var validate = component_interceptor.InterceptorComponent
	// TEMPLATE
	engine.GET("/v1/hello/hello", validate.GreetInterceptor, service.Hello)
	engine.GET("/v1/hello/echo", service.Echo)
}
