package server

import (

	//"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/zdq0394/simpleci/simpleci/config"
)

func setRoutes(r *gin.Engine, conf *config.Config) {
	ciService := NewCIService(conf)
	r.GET("/ping", ciService.ping)
	r.GET("/authorize", ciService.authorizeHanlder)
	r.GET("/callback", ciService.callbackHanlder)
	r.GET("/callback/codegot", ciService.codeGotHandler)
	r.GET("/callback/accessgot", ciService.accessTokenGotHandler)
	r.GET("/repos/:owner/:repo/hooks", ciService.createHookHandler)
}
