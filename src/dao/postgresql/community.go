package postgresql

import (
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"srbbs/src/model"
	"srbbs/src/query"
	"srbbs/src/srlogger"
	"time"
)

func GetCommunityById(id String) (*model.Community, err) {
	u := query.Community
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	first, err := query.Community.WithContext(ctx).Where(u.CommunityID.Eq(id)).First()
	if err != nil {
		srlogger.Logger().Warn("get community by id err", zap.Error(err))
		return nil, err
	}
	return first, err
}
