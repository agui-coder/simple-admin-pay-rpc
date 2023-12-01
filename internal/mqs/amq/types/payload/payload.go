// Package payload defines all the payload structures used in tasks
package payload

type PayOrderNotifyReq struct {
	MerchantOrderId string `json:"merchantOrderId"`
	PayOrderId      uint64 `json:"payOrderId"`
}

const (
	SUCCESS             = 0
	NotifyTimeoutMillis = 120
)

var NotifyFrequency = []int{15, 15, 30, 180,
	1800, 1800, 1800, 3600}

type PayOrderNotifyResp struct {
	// Error code | 错误代码
	Code int `json:"code"`
}
