package redis

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"srbbs/src/enums"
	"srbbs/src/model"
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
		"post:id":  postId,
		"user:id":  userId,
		"time":     now,
		"votes":    1,
		"comments": 0,
	}

	ctx := context.Background()
	pipeline := client.TxPipeline()

	// 默认帖主给自己点赞
	pipeline.ZAdd(ctx, votedKey, redis.Z{
		Score:  1,
		Member: userId,
	})
	pipeline.Expire(ctx, votedKey, 6*enums.OneMonthInSeconds*time.Second)

	// 文章 hash
	pipeline.HMSet(ctx, enums.KeyPostInfoHashPrefix+strconv.Itoa(int(postId)), postInfo)
	// 分数 zset
	pipeline.ZAdd(ctx, enums.KeyPostScoreZSet, redis.Z{
		Score:  enums.VoteScore,
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

// getIDsFormKey 按照分数从大到小的顺序查询指定数量的元素
func getIDsFormKey(key string, page, size int64) ([]string, error) {
	ctx := context.Background()

	start := (page - 1) * size
	end := start + size - 1
	return client.ZRevRange(ctx, key, start, end).Result()
}

// GetPostIDsInOrder 根据时间或分数排序
func GetPostIDsInOrder(p *model.C2SGetPostParam) ([]string, error) {
	// 从redis获取id
	// 1.根据用户请求中携带的order参数确定要查询的redis key
	key := enums.KeyPostTimeZSet     // 默认是时间
	if p.Order == enums.OrderScore { // 按照分数请求
		key = enums.KeyPostScoreZSet
	}
	// 2.确定查询的索引起始点
	return getIDsFormKey(key, p.Page, p.Size)
}

func GetPostVoteData(ids []string) (data []int64, err error) {
	ctx := context.Background()

	data = make([]int64, 0, len(ids))
	for _, id := range ids {
		key := enums.KeyPostVotedZSetPrefix + id
		// 查找key中分数是1的元素数量 -> 统计每篇帖子的赞成票的数量
		v := client.ZCount(ctx, key, "1", "1").Val()
		data = append(data, v)
	}
	return
}

func GetCommunityPostIDsInOrder(p *model.C2SGetPostParam) (ids []string, err error) {
	ctx := context.Background()

	// 1.根据用户请求中携带的order参数确定要查询的redis key
	orderkey := enums.KeyPostTimeZSet // 默认是时间
	if p.Order == enums.OrderScore {  // 按照分数请求
		orderkey = enums.KeyPostScoreZSet
	}

	// 针对新的zset 按之前的逻辑取数据

	// 社区的key
	cKey := enums.KeyCommunityPostSetPrefix + strconv.Itoa(int(p.CommunityID))

	key := orderkey + strconv.Itoa(int(p.CommunityID))
	if client.Exists(ctx, key).Val() < 1 {
		// 不存在，需要计算
		pipeline := client.Pipeline()

		// 把分区的帖子set与帖子分数的zset生成一个新的zset
		pipeline.ZInterStore(ctx, key, &redis.ZStore{
			Keys: []string{cKey, orderkey},
			//Weights:,
			Aggregate: "Max",
		})
		pipeline.Expire(ctx, key, 60*time.Second) // 设置超时时间
		_, err := pipeline.Exec(ctx)
		if err != nil {
			return nil, err
		}
	}
	// 存在的就直接根据key查询ids
	return getIDsFormKey(key, p.Page, p.Size)
}

func GetPostVoteNum(id int) int64 {
	ctx := context.Background()

	key := enums.KeyPostVotedZSetPrefix + strconv.Itoa(id)
	// 查找key中分数是1的元素数量 -> 统计每篇帖子的赞成票的数量
	return client.ZCount(ctx, key, "1", "1").Val()
}
