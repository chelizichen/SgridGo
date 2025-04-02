package service

import (
	"com_sgrid/src/components/constant"
	"com_sgrid/src/components/logger"
	"com_sgrid/src/utils"

	"github.com/gin-gonic/gin"
)
var log_service = logger.CreateLogger("service_index")

func Echo(c *gin.Context) {
	log_service.Info("invoke echo")
	utils.AbortWithSucc(c, constant.HELLO)
}

func Hello(c *gin.Context) {
	utils.AbortWithSucc(c, constant.WORLD)
}
