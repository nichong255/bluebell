package main

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/logger"
	"bluebell/settings"
	"fmt"

	"go.uber.org/zap"
)

//	A common template for Go Web development

func main() {
	// 1. load config file
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed, err:%v\n", err)
		return
	}
	// 2. init log
	if err := logger.Init(); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	zap.L().Debug("logger init success....")
	// 3. init mysql
	if err := mysql.Init(); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	// 4. init redis
	if err := redis.Init(); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	// 5. register router
	// 6. run service(smoothly shut down)
}
