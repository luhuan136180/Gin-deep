package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

var rdb *redis.Client

func initClirnt() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		PoolSize: 100,
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func redisExample() {
	err := rdb.Set("score", 100, 0).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}

	val, err := rdb.Get("score").Result()
	if err != nil {
		fmt.Printf("get score failed, err:%v\n", err)
		return
	}
	fmt.Println("score", val)

	val2, err := rdb.Get("name").Result()
	if err == redis.Nil {
		fmt.Println("name does not exist")
	} else if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
		return
	} else {
		fmt.Println("name", val2)
	}

}

//
func RedisListDemo() {
	//头插法创建List
	err := rdb.LPush("city", "beijing", "chengdu", "chongqing").Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}

	//尾插发
	err = rdb.RPush("city2", "tianjing", "海口", "昆明").Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}

	//遍历city
	city, err := rdb.LRange("city", 0, -1).Result()
	if err == redis.Nil { //先判断是否为空
		fmt.Println("city does not exist")
	} else if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
		return
	} else {
		fmt.Println("city", city)
	}

	//
	num, err := rdb.LLen("city").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("city的数量", num)

	//从右弹出元素
	val1, err := rdb.RPop("city").Result()
	if err == redis.Nil { //先判断是否为空
		fmt.Println("city does not exist")
	} else if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
		return
	} else {
		fmt.Println("city", val1)
	}
	//
	num, err = rdb.LLen("city").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("city的数量", num)

}

func RedisHashDemo() {
	err := rdb.HSet("user", "hoby", "足球").Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}
	v1, err := rdb.HGet("user", "name").Result()
	if err == redis.Nil { //先判断是否为空
		fmt.Println("city does not exist")
	} else if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
		return
	} else {
		fmt.Println("user name", v1)
	}
	v2 := rdb.HGetAll("user")
	fmt.Println("user:", v2)
}

func redisExample2() {
	zsetKey := "language_rank"
	language := []redis.Z{
		redis.Z{Score: 90.0, Member: "golang"},
		redis.Z{Score: 98.0, Member: "Java"},
		redis.Z{Score: 95.0, Member: "Python"},
		redis.Z{Score: 97.0, Member: "JavaScript"},
		redis.Z{Score: 99.0, Member: "C/C++"},
	}

	num, err := rdb.ZAdd(zsetKey, language...).Result()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Printf("zadd %d succ.\n", num)
	// 把Golang的分数加10
	newScore, err := rdb.ZIncrBy(zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)

	// 取分数最高的3个
	ret, err := rdb.ZRevRangeWithScores(zsetKey, 0, 2).Result()
	if err != nil {
		fmt.Printf("zrevrange failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

	// 取95~100分的
	op := redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = rdb.ZRangeByScoreWithScores(zsetKey, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}

func main() {
	if err := initClirnt(); err != nil {
		fmt.Printf("init redis client failed,err:%v\n", err)
		return
	}
	fmt.Println("Connect redis success....")

	//redisExample()
	//RedisListDemo()

	//RedisHashDemo()
	redisExample2()
}
