package Mainmain

import (
	"GPTManager/GlobalVar"
	"GPTManager/Main/Api"
	"GPTManager/Tool"
	"github.com/gin-gonic/gin"
)

func MainServer() {

	r := gin.Default()

	r.Use(Tool.Cors())
	r.Static("/html", "./HTML")
	r.POST("/Setprompt", Api.Setprompt)
	r.GET("/LocalIpPort", Api.LocalIpPort)
	r.POST("/CreateApi", Api.CreateApi)
	r.GET("/Init", Api.Init)
	r.POST("/SetSystem", Api.SetSystem)
	r.GET("/GetSystem", Api.GetSystem)
	r.POST("/StartCHAT", Api.StartCHAT)
	r.POST("/AppendApiKey", Api.AppendApiKey)
	r.GET("/AllocateAPIKEY", Api.AllocateAPIKEY)
	r.Run(":" + Tool.IntToStr(GlobalVar.Main_Port))
}
