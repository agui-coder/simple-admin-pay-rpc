package payment

import (
	"context"
	"fmt"
	"github.com/agui-coder/simple-admin-pay-rpc/payment/ali"
	"github.com/agui-coder/simple-admin-pay-rpc/payment/model"
	"github.com/agui-coder/simple-admin-pay-rpc/payment/weixin"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/alipay/cert"
	"log"
	"os"
	"strconv"
	"testing"
	"time"
)

var (
	factory = NewFactory()
)

func TestCreateWxPayClient(t *testing.T) {
	err := factory.CreateOrUpdatePayClient(model.WxPub, buildWxClientConfig())
	if err != nil {
		log.Fatal(err.Error())
	}
	client, err := factory.GetClient(model.WxPub)
	if err != nil {
		log.Fatal(err.Error())
	}
	order, err := client.UnifiedOrder(context.Background(), model.WxPub, buildPayOrderUnifiedReq())
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Print(&order)
}

func TestCreateAliPayClient(t *testing.T) {
	err := factory.CreateOrUpdatePayClient(model.AlipayPc, buildAliClientConfig())
	if err != nil {
		log.Fatal(err.Error())
	}
	client, err := factory.GetClient(model.AlipayPc)
	if err != nil {
		log.Fatal(err.Error())
	}
	resp, err := client.UnifiedOrder(context.Background(), model.AlipayPc, buildPayOrderUnifiedReq())
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Print(resp)
}

func buildPayOrderUnifiedReq() model.OrderUnifiedReq {
	m := make(map[string]string)
	// TODO openid
	m["openid"] = ""
	unifiedReq := model.OrderUnifiedReq{
		Price:         123,
		Subject:       "IPhone 13",
		Body:          "biubiubiu",
		OutTradeNo:    strconv.FormatInt(time.Now().UnixNano(), 10),
		UserIp:        "127.0.0.1",
		NotifyUrl:     "http://127.0.0.1:9107",
		ChannelExtras: m,
	}
	return unifiedReq
}

func buildWxClientConfig() weixin.ClientConfig {
	privateKeyFile, err := os.ReadFile("apiclient_key.pem")
	if err != nil {
		log.Fatal(err.Error())
	}
	//TODO
	return weixin.ClientConfig{
		AppId:             "",
		MchId:             "",
		PrivateKeyContent: privateKeyFile,
		SerialNumber:      "",
		ApiV3Key:          "",
	}
}

func buildAliClientConfig() ali.ClientConfig {
	return ali.ClientConfig{
		AppId:                   "代填",
		SignType:                alipay.RSA2,
		PrivateKey:              cert.PrivateKey,
		AppPublicContent:        cert.AppPublicContent,
		AlipayPublicContentRSA2: cert.AlipayPublicContentRSA2,
		AlipayRootContent:       cert.AlipayRootContent,
	}
}
