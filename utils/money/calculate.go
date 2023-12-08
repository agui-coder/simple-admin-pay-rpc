package money

import (
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"math"
)

// CalculateRatePriceInternal 计算百分比金额的内部函数
func CalculateRatePriceInternal(price int32, rate float64) (int32, error) {
	if rate == 0 {
		return 0, nil
	}
	if rate == 100 {
		return price, nil
	}

	result := float64(price) * rate / 100
	if result > float64(math.MaxInt32) {
		logx.Error("Result out of int32 range")
		return 0, errorx.NewInvalidArgumentError("Result out of int32 range")
	}

	return int32(result), nil
}
