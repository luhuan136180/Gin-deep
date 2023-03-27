package redis

const (
	Prefix                 = "bluebell:"
	KeyPostTimeZSet        = "post:time"
	KeyPostScoreZSet       = "post:score"
	KeyPostVotedZSetPrefix = "post:voted:"
	KeyCommunitySetPF      = "community:"
)

func GetRedisKey(key string) string {
	return Prefix + key
}
