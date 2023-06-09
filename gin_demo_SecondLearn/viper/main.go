package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//注意：viper将配置文件的内容解析到结构体中时，标签要用：mapstructure
type Config struct {
	Part        int    `mapstructure:"part"`
	Version     string `mapstructure:"version"`
	MySQLConfig `mapstructure:"mysql"`
}
type MySQLConfig struct {
	Host   string `mapstructure:"host"`
	Dbname string `mapstructure:"dbname"`
	Port   int    `mapstructure:"port"`
}

func main() {
	viper.SetDefault("fileDir", "./")

	//读取配置文件
	viper.SetConfigFile("./config.yaml") // 指定配置文件路径
	//viper.SetConfigName("config")         // 配置文件名称(无扩展名)
	//viper.SetConfigType("yaml")           // 如果配置文件的名称中没有扩展名，则需要配置此项
	//viper.AddConfigPath("/etc/appname/")  // 查找配置文件所在的路径
	//viper.AddConfigPath("$HOME/.appname") // 多次调用以添加多个搜索路径
	//viper.AddConfigPath(".")              // 还可以在工作目录中查找配置

	//
	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {             // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	//if err := viper.ReadInConfig(); err != nil {
	//	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
	//		// 配置文件未找到错误；如果需要可以忽略
	//	} else {
	//		// 配置文件被找到，但产生了另外的错误
	//	}
	//}

	//	监控并重新读取配置文件
	viper.WatchConfig()
	//当配置变化之后，调用一个回调函数
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
	})

	var c Config

	if err := viper.Unmarshal(&c); err != nil {
		fmt.Printf("viper.Unmarshal failed,err:%v\n", err)
		return
	}
	fmt.Printf("c:%#v\n", c)

	//r := gin.Default()
	//r.GET("/version", func(c *gin.Context) {
	//	c.String(http.StatusOK, viper.GetString("version"))
	//})
	//r.Run()
}
