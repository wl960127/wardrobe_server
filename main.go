package main

import (
	"fmt"
	"net/http"

	router "wardrobe_server/routers"
	parse_config "wardrobe_server/pkg/app/parseConfig"

)

func main() {

	// 注册接口
	router := router.InitRouter()

	
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", parse_config.ServerSetting.HTTPPort),
		Handler:        router,
		ReadTimeout:    parse_config.ServerSetting.ReadTimeout,
		WriteTimeout:   parse_config.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	// http.Handle("/ws", v1.WsManager)




	s.ListenAndServe()
}
