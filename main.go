package main

import (
	"gin-user/config"
	"gin-user/router"
)

func main() {
	r := router.InitRouter()
	panic(r.Run(":"+config.Conf.Server.Port))
}

