package snowFlake

import (
	"github.com/bwmarrin/snowflake"
	"time"
)

var node *snowflake.Node

func Init(startTime string, machineID int64) (err error) {
	var st time.Time //时间因子
	//对输入的起始时间进行规格化处理
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return

	}
	snowflake.Epoch = st.UnixNano() / 1000000
	//用机器号创建node
	node, err = snowflake.NewNode(machineID)
	return
}

func GenID() int64 {
	return node.Generate().Int64()
}
