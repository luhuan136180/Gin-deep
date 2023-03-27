package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

//var rdb *redis.Client
//
//func initClirnt() (err error) {
//	rdb = redis.NewClient(&redis.Options{
//		Addr:     "localhost:6379",
//		Password: "",
//		DB:       0,
//		PoolSize: 100,
//	})
//
//	_, err = rdb.Ping().Result()
//	if err != nil {
//		return err
//	}
//	return nil
//}
func WatchDemo() {
	key := "watch_count"
	if err := rdb.Watch(func(tx *redis.Tx) error {
		n, err := tx.Get(key).Int()
		if err != nil && err != redis.Nil {
			return err
		}

		_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
			time.Sleep(time.Second * 6) // 此时在另外一个终端窗口中使用set命令修改watch_count的值
			pipe.Set(key, n+1, 0)
			return nil
		})
		return err
	}, key); err != nil {
		fmt.Printf("key发生变化，更新失败。err:%v\n", err)
	} else {
		fmt.Println("更新成功")
	}

}

func main() {

	if err := initClirnt(); err != nil {
		fmt.Printf("init redis client failed,err:%v\n", err)
		return
	}
	fmt.Println("Connect redis success....")

	WatchDemo()

}
