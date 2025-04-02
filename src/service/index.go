package service

import (
	"com_sgrid/src/components/constant"
	"com_sgrid/src/utils"

	"github.com/gin-gonic/gin"
)

func Echo(c *gin.Context) {
	utils.AbortWithSucc(c, constant.HELLO)
}

func Hello(c *gin.Context) {
	utils.AbortWithSucc(c, constant.WORLD)
}
