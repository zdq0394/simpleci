package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/zdq0394/simpleci/simpleci/config"
)

const (
	defaultListenAddr = ":8080"
)

var AccessToken string

func Start(configPath string) {
	conf, err := config.LoadConfig(configPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	r := gin.Default()
	// set routes
	setRoutes(r, &conf)
	// start
	startAt(r, defaultListenAddr)
}

func startAt(r *gin.Engine, addrs ...string) {
	switch len(addrs) {
	case 0:
		go func() {
			r.Run(defaultListenAddr)
		}()
	default:
		for _, addr := range addrs {
			t_add := addr
			go func() {
				r.Run(t_add)
			}()
		}
	}
	select {}

}
