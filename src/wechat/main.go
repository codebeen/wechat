package main

import (
	"wechat/utils"
	"wechat/router"
)

func main() {
	if err := utils.InitConfig(); err != nil {
		panic(err)
	}

	// router
	r := router.InitRouter()
	r.Run(utils.Conf.App.Port)
}
