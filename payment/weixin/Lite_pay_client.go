package weixin

// litePayStrategy 结构体继承了 PubPayStrategy 接口 由于公众号和小程序的微信支付逻辑一致，所以直接进行继承
type litePayStrategy struct {
	pubPayStrategy
}
