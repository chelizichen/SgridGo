package main

import (
	"com_sgrid/src/components"
	component_cache "com_sgrid/src/components/cache"
	"com_sgrid/src/components/constant"
	component_db "com_sgrid/src/components/db"
	"com_sgrid/src/components/logger"
	"com_sgrid/src/components/middleware"
	"com_sgrid/src/router"
	"com_sgrid/src/schedule"
	"fmt"
	"os"
)

var log_app = logger.CreateLogger("app")

func init() {
	components.LoadComponents()
	components.Gin_Server.Use(middleware.Cors())
	router.LoadRouter(components.Gin_Server)
	component_db.LoadDB(components.Sgrid_Conf)
	component_cache.LoadCache(components.Sgrid_Conf)
	defer schedule.InitSchedule()
}

func main() {
	var port = fmt.Sprintf(":%s", os.Getenv(constant.ENV_TARGET_PORT))
	if port == ":" {
		port = fmt.Sprintf(":%v", components.Sgrid_Conf.Server.Port)
	}
	log_app.Info("Starting server on port :", port)
	components.Gin_Server.Run(port)
}
