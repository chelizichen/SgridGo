package components

import (
	"com_sgrid/src/framework/config"
	"com_sgrid/src/framework/public"
	"fmt"

	"github.com/gin-gonic/gin"
)

var Gin_Server *gin.Engine
var Sgrid_Conf *config.SgridConf

func LoadComponents() {
	Gin_Server = gin.Default()
	conf, err := public.NewConfig()
	if err != nil {
		panic("error to New Config")
	}
	Sgrid_Conf = conf
	fmt.Println("load conf :: ", conf)
}
