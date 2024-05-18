package enums

const (
	KeyPostInfoHashPrefix = "srbbbs:post:"
	KeyPostTimeZSet       = "srbbbs:post:time"  // zset;帖子及发帖时间定义
	KeyPostScoreZSet      = "srbbbs:post:score" // zset;帖子及投票分数定义
	//KeyPostVotedUpSetPrefix   = "srbbbs:post:voted:down:"
	//KeyPostVotedDownSetPrefix = "srbbbs:post:voted:up:"
	KeyPostVotedZSetPrefix    = "srbbbs:post:voted:" // zSet;记录用户及投票类型;参数是post_id
	KeyCommunityPostSetPrefix = "srbbbs:community:"  // set保存每个分区下帖子的id
)

const (
	OrderTime  = "time"
	OrderScore = "score"
)
