package component_interceptor

import (
	"github.com/gin-gonic/gin"
)

type componentInterceptor struct{}

var InterceptorComponent = new(componentInterceptor)

func (i *componentInterceptor) GreetInterceptor(c *gin.Context) {
	c.Next()
}
