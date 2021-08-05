package main

import (
	"todo/global"
	"todo/initialize"
)

// @title TodoList API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api/
func main() {

	cfg := initialize.InitConfig(global.CfgFilePath)
	app := initialize.NewApp(&cfg)
	app.Run()

}
