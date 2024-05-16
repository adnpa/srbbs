package postgresql

import (
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"srbbs/src/model"
	"srbbs/src/query"
	"srbbs/src/srlogger"
	"time"
)

func CreatePost(post *model.Post) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = query.Post.WithContext(ctx).Create(post)
	if err != nil {
		srlogger.Logger().Error("Error creating post", zap.Error(err))
	}
	return
}
