// Package payload defines all the payload structures used in tasks
package payload

import "github.com/agui-coder/simple-admin-pay-rpc/pay"

type PayOrderNotifyReq struct {
	MerchantOrderId string `json:"merchantOrderId"`
	PayOrderId      uint64 `json:"payOrderId"`
}

type PayRefundNotifyReq struct {
	MerchantOrderId string `json:"merchantOrderId"`
	PayRefundId     uint64 `json:"payRefundId"`
}

const (
	SUCCESS             = pay.PayStatus_PAY_SUCCESS
	NotifyTimeoutMillis = 120
)

var NotifyFrequency = []int{15, 15, 30, 180,
	1800, 1800, 1800, 3600}

type PayOrderNotifyResp struct {
	// Error code | 错误代码
	Code int `json:"code"`
	// Error message | 错误消息
	Msg string `json:"msg"`
}
