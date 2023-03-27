package main

import (
	"bluebell0002/controller"
	"bluebell0002/dao/mysql"
	"bluebell0002/dao/redis"
	"bluebell0002/logger"
	"bluebell0002/pkg/snowFlake"
	"bluebell0002/router"
	"bluebell0002/settings"
	"fmt"
	"go.uber.org/zap"
)

func main() {
	//1.加载配置——（远程/配置文件中）
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed,err:%#v\n", err)
		return
	}
	//2.初始化日志
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed,err:%#v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success。。。")
	//3.初始化Mysql
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed,err:%#v\n", err)
		return
	}
	defer mysql.DB.Close()

	//4.初始化Redis
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed,err:%#v\n", err)
		return
	}
	defer redis.Close()
	//初始化雪花算法ID生成器
	if err := snowFlake.Init(settings.Conf.StartTime, settings.Conf.MachineId); err != nil {
		fmt.Printf("init snowflake failed,err:%#v\n", err)
		return
	}

	//初始化gin框架内置的校验器使用的翻译器
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init transfer failed,err:%#v\n", err)
		return
	}
	//5.注册路由

	r := router.SetUpRouter(settings.Conf.Mode)
	err := r.Run(fmt.Sprintf(":%d", settings.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
