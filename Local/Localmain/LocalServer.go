package Localmain

import (
	"GPTManager/GlobalVar"
	"GPTManager/Local/Api"
	"GPTManager/Tool"
	"github.com/gin-gonic/gin"
)

//	func login(w http.ResponseWriter, r *http.Request) {
//		log.Printf("login")
//		w.Write([]byte(string("login")))
//	}
//
// localMain
func LocalServer() {
	Init()
	r := gin.Default()
	r.Use(Tool.Cors())
	r.POST("/GetApiKey", Api.GetApiKey)
	r.POST("/CHAT", Api.CHAT)
	r.Run(":" + Tool.IntToStr(GlobalVar.Port))
}
