package payno

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	PayNo          = "PAY_NO:"
	OrderNoPrefix  = "P"
	RefundNoPrefix = "R"
)

// Generate 生成支付编号
func Generate(ctx context.Context, db redis.UniversalClient, prefix string) string {
	// 构建没有前缀的编号
	noPrefix := prefix + time.Now().Format("20060102150405")
	key := PayNo + noPrefix
	incr := db.Incr(ctx, key)
	db.Expire(ctx, key, 60)
	return noPrefix + incr.String()
}
