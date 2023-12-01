package money

import (
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"math"
	"math/big"
)

// CalculateRatePriceInternal 计算百分比金额的内部函数
func CalculateRatePriceInternal(price int32, rate float64) (int32, error) {
	priceBigInt := big.NewFloat(float64(price))

	rateFloat := new(big.Float).SetFloat64(rate)

	resultFloat := new(big.Float)
	resultFloat.Mul(rateFloat, priceBigInt)
	resultFloat.Quo(resultFloat, big.NewFloat(100))

	result := new(big.Int)
	resultFloat.Int(result)
	int64Value := result.Int64()
	if int64Value > math.MaxInt32 {
		logx.Error("big.Int value out of int32 range")
		return 0, errorx.NewInvalidArgumentError("big.Int value out of int32 range")
	}
	int32Value := int32(int64Value)

	return int32Value, nil
}
