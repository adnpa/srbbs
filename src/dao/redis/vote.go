package redis

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"srbbs/src/enums"
	"srbbs/src/srlogger"
	"strconv"
	"time"
)

// 赞/踩功能

func CreatePost(postId, userId, communityId int64, title, summary string) error {
	now := float64(time.Now().Unix())
	votedKey := enums.KeyPostVotedZSetPrefix + strconv.Itoa(int(postId))
	communityKey := enums.KeyCommunityPostSetPrefix + strconv.Itoa(int(communityId))
	postInfo := map[string]interface{}{
		"title":    title,
		"summary":  summary,
		"post:id":  postID,
		"user:id":  userID,
		"time":     now,
		"votes":    1,
		"comments": 0,
	}

	ctx := context.Background()
	pipeline := client.TxPipeline()

	pipeline.ZAdd(ctx, votedKey, redis.Z{
		Score:  1,
		Member: userId,
	})
	pipeline.Expire(ctx, votedKey, 6*enums.OneMonthInSeconds*time.Second)
	// 文章 hash
	pipeline.HMSet(ctx, enums.KeyPostInfoHashPrefix+strconv.Itoa(int(postId)), postInfo)
	// 分数 zset
	pipeline.ZAdd(ctx, enums.KeyPostScoreZSet, redis.Z{
		Score:  now + enums.VoteScore,
		Member: postId,
	})
	//	时间 zset
	pipeline.ZAdd(ctx, enums.KeyPostTimeZSet, redis.Z{
		Score:  now,
		Member: postId,
	})

	// 社区 set
	pipeline.SAdd(ctx, communityKey, postId)
	_, err := pipeline.Exec(ctx)
	if err != nil {
		srlogger.Logger().Warn("CreatePost", zap.Error(err))
	}
	return err
}
