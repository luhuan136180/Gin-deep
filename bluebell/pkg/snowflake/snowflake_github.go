package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"time"
)

//启动全局node
var node *snowflake.Node

func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	//st.UnixNano()返回的是单位纳秒的unix时间,
	//snowflake.Epoch是单位为毫秒的时间,记录的是雪花算法的id计算的起始时间戳，可以自定义
	//**自设时需要单位转换
	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineID)
	return
}

//生成int64的雪花算法uid值

func GenID() int64 {
	return node.Generate().Int64()
}

//func main() {
//	if err := Init("2020-07-01", 1); err != nil {
//		fmt.Printf("init failed, err:%v\n", err)
//		return
//	}
//	id := GenID()
//	fmt.Println(id)
//}
