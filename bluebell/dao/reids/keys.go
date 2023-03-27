package redis

//reids key

// redis key 使用命名空间的方式，方便查询和拆分

const (
	Prefix                 = "bluebell:"           // 项目key前缀
	KeyPostTImeZSet        = "bluebell:post:time"  //zset:以帖子发布的时间
	KeyPostScoreZSet       = "bluebell:post:score" //zset:以帖子投票的分数
	KeyPostVotedZSetPrefix = "post:voted:"         //zset:记录用户的id及投票的类型，需要的参数是帖子的id

	KeyCommunitySetPF = "community:" // set;保存每个分区下帖子的id
)

// 给redis key加上前缀
func getRedisKey(key string) string {
	return Prefix + key
}
